package utils

const WssUrl = "livetiming.formula1.com"
const SignalrUrl = "livetiming.formula1.com/signalr"

// const SignalrHub = "Streaming";
// const SignalrHubParsed = `%5B%7B%22name%22%3A%22Streaming%22%7D%5D`
const SignalrHubParsed = `[{ "name": "Streaming" }]`
const SignalrSubscribe = `{
    "H": "Streaming",
    "M": "Subscribe",
    "A": [[
        "Heartbeat",
        "CarData.z",
        "Position.z",
        "ExtrapolatedClock",
        "TopThree",
        "RcmSeries",
        "TimingStats",
        "TimingAppData",
        "WeatherData",
        "TrackStatus",
        "DriverList",
        "RaceControlMessages",
        "SessionInfo",
        "SessionData",
        "LapCount",
        "TimingData",
        "TeamRadio",
        "PitLaneTimeCollection",
        "ChampionshipPrediction"
    ]],
    "I": 1,
}`
