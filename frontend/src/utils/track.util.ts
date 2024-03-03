type StatusMessage = {
  message: string;
  color: string;
};

type MessageMap = {
  [key: string]: StatusMessage;
};

export const getTrackStatusMessage = (
  statusCode: number | undefined
): StatusMessage | null => {
  const messageMap: MessageMap = {
    1: { message: "Track Clear", color: "bg-success" },
    2: { message: "Yellow Flag", color: "bg-warning" },
    3: { message: "Flag", color: "bg-warning" },
    4: { message: "Safety Car", color: "bg-warning" },
    5: { message: "Red Flag", color: "bg-error" },
    6: { message: "VSC Deployed", color: "bg-warning" },
    7: { message: "VSC Ending", color: "bg-warning" },
  };

  return statusCode ? messageMap[statusCode] ?? messageMap[0] : null;
};
