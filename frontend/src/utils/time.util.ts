import moment from "moment";

export const getTimeColor = (fastest: boolean, pb: boolean) => {
  if (fastest) return "text-violet-600";
  else if (pb) return "text-success";
  return "";
};

type UtcObject = { utc: string };

export const sortUtc = (a: UtcObject, b: UtcObject) => {
  return moment.utc(b.utc).diff(moment.utc(a.utc));
};
