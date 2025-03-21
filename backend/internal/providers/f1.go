package provider

import (
	"context"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"sync"
	"time"

	"github.com/f1gopher/f1gopherlib/Messages"
	"github.com/f1gopher/f1gopherlib/connection"
	"github.com/f1gopher/f1gopherlib/f1log"
	"github.com/f1gopher/f1gopherlib/flowControl"
	"github.com/f1gopher/f1gopherlib/internal/parser"
)

type F1Lib interface {
	Name() string
	Session() Messages.SessionType
	CircuitTimezone() *time.Location
	SessionStart() time.Time
	Track() string
	TrackYear() int
	TimeLostInPitlane() time.Duration

	Weather() <-chan Messages.Weather
	RaceControlMessages() <-chan Messages.RaceControlMessage
	Timing() <-chan Messages.Timing
	Event() <-chan Messages.Event
	Telemetry() <-chan Messages.Telemetry
	Location() <-chan Messages.Location
	Time() <-chan Messages.EventTime
	Radio() <-chan Messages.Radio
	Drivers() <-chan Messages.Drivers

	Data() any

	SelectTelemetrySources(drivers []int)

	IncrementLap()
	IncrementTime(duration time.Duration)
	SkipToSessionStart()
	TogglePause()
	IsPaused() bool

	Close()
}

type f1lib struct {
	archive string

	session           Messages.SessionType
	name              string
	timezone          *time.Location
	sessionStart      time.Time
	track             string
	trackYear         int
	timeLostInPitlane time.Duration

	connection   connection.Connection
	dataHandler  *parser.Parser
	replayTiming flowControl.Flow

	weather             chan Messages.Weather
	raceControlMessages chan Messages.RaceControlMessage
	timing              chan Messages.Timing
	event               chan Messages.Event
	telemetry           chan Messages.Telemetry
	location            chan Messages.Location
	eventTime           chan Messages.EventTime
	radio               chan Messages.Radio
	drivers             chan Messages.Drivers

	ctxShutdown context.CancelFunc
	ctx         context.Context
	wg          sync.WaitGroup
}

// TODO - do we ned different sizes for straight through and realtime flow control? Straight through will need
// bigger buffers but realtime only needs to hold abotu 5 mins of data
const weatherChannelSize = 100
const rcmChannelSize = 100
const timingChannelSize = 10000
const eventChannelSize = 1000
const telemetryChannelSize = 1000
const locationChannelSize = 1000
const eventTimeChannelSize = 10
const radioChannelSize = 100
const driversChannelSize = 100

var f1Log = f1log.CreateLog()

func SetLogOutput(w io.Writer) {
	f1Log.SetLogOutput(w)
}

func CreateRaceEvent(
	country string,
	raceTime time.Time,
	eventTime time.Time,
	sessionType Messages.SessionType,
	name string,
	track string,
	trackYearCreated int,
	pitlaneTime time.Duration,
	urlName string,
	timezone string) *RaceEvent {

	var sessionName string
	switch sessionType {
	case Messages.Practice1Session:
		sessionName = "Practice_1"
	case Messages.Practice2Session:
		sessionName = "Practice_2"
	case Messages.Practice3Session:
		sessionName = "Practice_3"
	case Messages.QualifyingSession:
		sessionName = "Qualifying"
	case Messages.SprintSession:
		sessionName = "Sprint"
	case Messages.RaceSession:
		sessionName = "Race"
	case Messages.PreSeasonSession:
		sessionName = "Test"
	default:
		panic("Unhandled session type: " + sessionType.String())
	}

	urlName = fmt.Sprintf(
		"https://livetiming.formula1.com/static/%d/%d-%02d-%02d_%s_Grand_Prix/%d-%02d-%02d_%s/",
		raceTime.Year(),
		raceTime.Year(),
		raceTime.Month(),
		raceTime.Day(),
		urlName,
		eventTime.Year(),
		eventTime.Month(),
		eventTime.Day(),
		sessionName)

	return &RaceEvent{
		Country:           country,
		RaceTime:          raceTime,
		EventTime:         eventTime,
		Type:              sessionType,
		Name:              name,
		timezone:          timezone,
		TrackName:         track,
		TrackYearCreated:  trackYearCreated,
		urlName:           urlName,
		TimeLostInPitlane: pitlaneTime,
	}
}

type RaceEvent struct {
	Country           string
	RaceTime          time.Time
	EventTime         time.Time
	Type              Messages.SessionType
	Name              string
	timezone          string
	TrackName         string
	TrackYearCreated  int
	TimeLostInPitlane time.Duration

	// TODO - add duration

	urlName string
}

func (r *RaceEvent) Timezone() *time.Location {
	tz, _ := time.LoadLocation(r.timezone)
	return tz
}

func (r *RaceEvent) string() string {
	return fmt.Sprintf("%s - %s", r.Name, r.Type.String())
}

func (r *RaceEvent) Url() string {
	return r.urlName
}

func CreateLive(requestedData parser.DataSource, archive string, cache string) (F1Lib, error) {

	// TODO - validate path
	// TODO - create archive folder

	currentEvent, exists := liveEvent()

	// No event happening or about to happen so nothing we can do
	if !exists {
		return nil, errors.New("no live event currently happening")
	}

	f1Log.Infof("Creating live session for: %v", currentEvent.string())

	data := f1lib{
		weather:             make(chan Messages.Weather, weatherChannelSize),
		raceControlMessages: make(chan Messages.RaceControlMessage, rcmChannelSize),
		timing:              make(chan Messages.Timing, timingChannelSize),
		event:               make(chan Messages.Event, eventChannelSize),
		telemetry:           make(chan Messages.Telemetry, telemetryChannelSize),
		location:            make(chan Messages.Location, locationChannelSize),
		eventTime:           make(chan Messages.EventTime, eventTimeChannelSize),
		radio:               make(chan Messages.Radio, radioChannelSize),
		drivers:             make(chan Messages.Drivers, driversChannelSize),

		archive:           archive,
		session:           currentEvent.Type,
		name:              currentEvent.Name,
		timezone:          currentEvent.Timezone(),
		sessionStart:      currentEvent.EventTime,
		track:             currentEvent.TrackName,
		trackYear:         currentEvent.TrackYearCreated,
		timeLostInPitlane: currentEvent.TimeLostInPitlane,
	}
	data.ctx, data.ctxShutdown = context.WithCancel(context.Background())

	err := data.connectLive(requestedData, archive, currentEvent, cache)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func CreateDebugReplay(
	requestedData parser.DataSource,
	replayFile string,
	dataFlow flowControl.FlowType) (F1Lib, error) {

	// TODO - read event info from debug file
	event := RaceEvent{}

	f1Log.Infof("Creating live replay session for: %v", event.string())

	data := f1lib{
		weather:             make(chan Messages.Weather, weatherChannelSize),
		raceControlMessages: make(chan Messages.RaceControlMessage, rcmChannelSize),
		timing:              make(chan Messages.Timing, timingChannelSize),
		event:               make(chan Messages.Event, eventChannelSize),
		telemetry:           make(chan Messages.Telemetry, telemetryChannelSize),
		location:            make(chan Messages.Location, locationChannelSize),
		eventTime:           make(chan Messages.EventTime, eventTimeChannelSize),
		radio:               make(chan Messages.Radio, radioChannelSize),
		drivers:             make(chan Messages.Drivers, driversChannelSize),
		session:             event.Type,
		name:                event.Name,
		timezone:            event.Timezone(),
		sessionStart:        event.EventTime,
		track:               event.TrackName,
		trackYear:           event.TrackYearCreated,
		timeLostInPitlane:   event.TimeLostInPitlane,
	}
	data.ctx, data.ctxShutdown = context.WithCancel(context.Background())

	err := data.connectDebugReplay(requestedData, replayFile, event, dataFlow)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func CreateReplay(
	requestedData parser.DataSource,
	event RaceEvent,
	cache string,
	dataFlow flowControl.FlowType) (F1Lib, error) {

	f1Log.Infof("Creating replay session for: %v", event.string())

	data := f1lib{
		weather:             make(chan Messages.Weather, weatherChannelSize),
		raceControlMessages: make(chan Messages.RaceControlMessage, rcmChannelSize),
		timing:              make(chan Messages.Timing, timingChannelSize),
		event:               make(chan Messages.Event, eventChannelSize),
		telemetry:           make(chan Messages.Telemetry, telemetryChannelSize),
		location:            make(chan Messages.Location, locationChannelSize),
		eventTime:           make(chan Messages.EventTime, eventTimeChannelSize),
		radio:               make(chan Messages.Radio, radioChannelSize),
		drivers:             make(chan Messages.Drivers, driversChannelSize),
		session:             event.Type,
		name:                event.Name,
		timezone:            event.Timezone(),
		sessionStart:        event.EventTime,
		track:               event.TrackName,
		trackYear:           event.TrackYearCreated,
		timeLostInPitlane:   event.TimeLostInPitlane,
	}
	data.ctx, data.ctxShutdown = context.WithCancel(context.Background())

	err := data.connectReplay(requestedData, event, cache, dataFlow)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (f *f1lib) connectLive(requestedData parser.DataSource, archiveFile string, event RaceEvent, cache string) error {

	cache = f.cachePath(cache, event)

	if len(archiveFile) == 0 {
		f.connection = connection.CreateLive(f.ctx, &f.wg, f1Log)
	} else {
		var connErr error
		f.connection, connErr = connection.CreateArchivingLive(f.ctx, &f.wg, f1Log, archiveFile)
		if connErr != nil {
			return connErr
		}
	}

	err, dataChannel := f.connection.Connect()
	if err != nil {
		return err
	}

	f.replayTiming = flowControl.CreateFlowControl(
		f.ctx,
		&f.wg,
		flowControl.Realtime,
		f.weather,
		f.raceControlMessages,
		f.timing,
		f.event,
		f.telemetry,
		f.location,
		f.eventTime,
		f.radio,
		f.drivers)

	assetStore := connection.CreateAssetStore(event.Url(), cache, f1Log)

	f.dataHandler = parser.Create(
		f.ctx,
		&f.wg,
		requestedData,
		dataChannel,
		f.replayTiming,
		assetStore,
		Messages.RaceSession,
		f1Log,
		event.Timezone())

	go f.dataHandler.Process()
	go f.replayTiming.Run()

	return nil
}

func (f *f1lib) connectDebugReplay(
	requestedData parser.DataSource,
	replayFile string,
	event RaceEvent,
	dataFlow flowControl.FlowType) error {

	f.connection = connection.CreateArchivedLive(f.ctx, &f.wg, f1Log, replayFile)
	err, dataChannel := f.connection.Connect()

	if err != nil {
		return err
	}

	f.replayTiming = flowControl.CreateFlowControl(
		f.ctx,
		&f.wg,
		dataFlow,
		f.weather,
		f.raceControlMessages,
		f.timing,
		f.event,
		f.telemetry,
		f.location,
		f.eventTime,
		f.radio,
		f.drivers)

	// Don't use a cache for debug replays because we don't always know the event yet to give it a useful folder name
	assetStore := connection.CreateAssetStore(event.Url(), "", f1Log)

	f.dataHandler = parser.Create(
		f.ctx,
		&f.wg,
		requestedData,
		dataChannel,
		f.replayTiming,
		assetStore,
		event.Type,
		f1Log,
		event.Timezone())

	go f.dataHandler.Process()
	go f.replayTiming.Run()

	return nil
}

func (f *f1lib) connectReplay(
	requestedData parser.DataSource,
	event RaceEvent,
	cache string,
	dataFlow flowControl.FlowType) error {

	url := event.Url()
	cache = f.cachePath(cache, event)

	f.connection = connection.CreateReplay(f.ctx, &f.wg, f1Log, url, event.Type, event.RaceTime.Year(), cache)
	err, dataChannel := f.connection.Connect()

	if err != nil {
		return err
	}

	f.replayTiming = flowControl.CreateFlowControl(
		f.ctx,
		&f.wg,
		dataFlow,
		f.weather,
		f.raceControlMessages,
		f.timing,
		f.event,
		f.telemetry,
		f.location,
		f.eventTime,
		f.radio,
		f.drivers)

	assetStore := connection.CreateAssetStore(event.Url(), cache, f1Log)

	f.dataHandler = parser.Create(
		f.ctx,
		&f.wg,
		requestedData,
		dataChannel,
		f.replayTiming,
		assetStore,
		event.Type,
		f1Log,
		event.Timezone())

	go f.dataHandler.Process()
	go f.replayTiming.Run()

	return nil
}

func (f *f1lib) cachePath(cache string, event RaceEvent) string {
	return filepath.Join(cache, fmt.Sprintf("%d", event.RaceTime.Year()), fmt.Sprintf("%s_%s", event.RaceTime.Format("2006-01-02"), event.Name), event.Type.String())
}

func (f *f1lib) Session() Messages.SessionType {
	return f.session
}

func (f *f1lib) Name() string {
	return f.name
}

func (f *f1lib) CircuitTimezone() *time.Location {
	return f.timezone
}

func (f *f1lib) SessionStart() time.Time {
	return f.sessionStart
}

func (f *f1lib) Track() string {
	return f.track
}

func (f *f1lib) TrackYear() int {
	return f.trackYear
}

func (f *f1lib) TimeLostInPitlane() time.Duration {
	return f.timeLostInPitlane
}

func (f *f1lib) Weather() <-chan Messages.Weather {
	return f.weather
}

func (f *f1lib) RaceControlMessages() <-chan Messages.RaceControlMessage {
	return f.raceControlMessages
}

func (f *f1lib) Timing() <-chan Messages.Timing {
	return f.timing
}

func (f *f1lib) Event() <-chan Messages.Event {
	return f.event
}

func (f *f1lib) Telemetry() <-chan Messages.Telemetry {
	return f.telemetry
}

func (f *f1lib) Location() <-chan Messages.Location {
	return f.location
}

func (f *f1lib) Time() <-chan Messages.EventTime {
	return f.eventTime
}

func (f *f1lib) Radio() <-chan Messages.Radio {
	return f.radio
}

func (f *f1lib) Drivers() <-chan Messages.Drivers {
	return f.drivers
}

func (f *f1lib) Data() any {
	return &f1lib{
		weather:           f.weather,
		session:           f.session,
		telemetry:         f.telemetry,
		event:             f.event,
		drivers:           f.drivers,
		timing:            f.timing,
		name:              f.name,
		timeLostInPitlane: f.timeLostInPitlane,
		radio:             f.radio,
		track:             f.track,
		sessionStart:      f.sessionStart,
		location:          f.location,
		eventTime:         f.eventTime,
	}
}

func (f *f1lib) SelectTelemetrySources(drivers []int) {
	f.dataHandler.SelectTelemetrySources(drivers)
}

func (f *f1lib) IncrementLap() {
	// Only makes sense for races
	if f.session == Messages.RaceSession || f.session == Messages.SprintSession {
		f.replayTiming.IncrementLap()
	}
}

func (f *f1lib) IncrementTime(duration time.Duration) {
	f.connection.IncrementTime(duration)
	f.replayTiming.IncrementTime(duration)
}

func (f *f1lib) SkipToSessionStart() {
	sessionStart := f.connection.JumpToStart()
	fmt.Println("sessionStart", sessionStart)
	if !sessionStart.IsZero() {
		f.replayTiming.SkipToSessionStart(sessionStart)
	}
}

func (f *f1lib) TogglePause() {
	f.replayTiming.TogglePause()
}

func (f *f1lib) IsPaused() bool {
	return f.replayTiming.IsPaused()
}

func (f *f1lib) Close() {
	f.name = ""
	f.track = ""

	f.ctxShutdown()
	f.wg.Wait()

	f.connection = nil
	f.dataHandler = nil

	close(f.weather)
	close(f.raceControlMessages)
	close(f.timing)
	close(f.event)
	close(f.telemetry)
	close(f.location)
	close(f.eventTime)
	close(f.radio)
	close(f.drivers)
}
