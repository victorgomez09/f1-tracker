import { config } from "../config";

type NegotiateResult = {
  Url: string;
  ConnectionToken: string;
  ConnectionId: string;
  KeepAliveTimeout: number;
  DisconnectTimeout: number;
  ConnectionTimeout: number;
  TryWebSockets: boolean;
  ProtocolVersion: string;
  TransportConnectTimeout: number;
  LongPollDelay: number;
};

export const negotiate = async (hub: string) => {
  if (!!config.testing)
    return {
      body: {
        Url: "string",
        ConnectionToken: "string",
        ConnectionId: "string",
        KeepAliveTimeout: 0,
        DisconnectTimeout: 0,
        ConnectionTimeout: 0,
        TryWebSockets: true,
        ProtocolVersion: "string",
        TransportConnectTimeout: 1,
        LongPollDelay: 1,
      },
      cookie: "",
    };

  const url = `${config.f1NegotiateUrl}/negotiate?connectionData=${hub}&clientProtocol=1.5`;
  const res = await fetch(url);

  const body: NegotiateResult = await res.json();

  return {
    body,
    cookie:
      res.headers.get("Set-Cookie") ?? res.headers.get("set-cookie") ?? "",
  };
};
