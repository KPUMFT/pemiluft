import moment from "moment";

export const dapilValues = [
  "TEKNIK INFORMATIKA",
  "TEKNIK ELEKTRO",
  "TEKNIK INDUSTRI",
  "SISTEM INFORMASI",
  "TEKNIK MESIN",
  "TEKNIK MEKATRONIKA",
];
export const startDate = new Date(2021, 0, 4, 8, 0, 0);
export const middleDate = new Date(2021, 0, 4, 10, 0, 0);
export const endDate = new Date(2021, 0, 4, 13, 0, 0);
export const wrongStartDate = new Date(2021, 0, 4, 13, 1, 0);
export const wrongEndDate = new Date(2021, 0, 4, 14, 0, 0);

export const formatDate = (v) =>
  moment(v).format("YYYY/MM/DD h:mm:ss A") + " GMT+7";

export const alphabetOnly = (v) => v.replace(/[^a-zA-Z]+/g, "");
export const last3Letters = (v) => v.substr(v.length - 3);
export const gubernurValues = [
  "Calon Gubernur nomor urut 1",
  "Calon Gubernur nomor urut 2",
  "Calon Gubernur nomor urut 3",
  "Calon Gubernur nomor urut 4",
];

export const dpmValues = [
  "Calon DPM nomor urut 1",
  "Calon DPM nomor urut 2",
  "Calon DPM nomor urut 3",
];

export const calonDapilValues = [
  "Calon nomor urut 1",
  "Calon nomor urut 2",
  "Calon nomor urut 3",
  "Calon nomor urut 4",
];

export const zigZagMod = (i, divisor) => {
  const halfDivisor = divisor / 2;
  const mod = (i + 1) % divisor;
  const logic = mod === 0 ? false : mod <= halfDivisor;
  return logic;
};
