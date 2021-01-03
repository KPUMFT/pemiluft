import fs from "fs";
import stringify from "csv-stringify/lib/sync.js";

import faker from "./fakerInit.js";
import {
  startDate,
  middleDate,
  wrongStartDate,
  wrongEndDate,
  formatDate,
  gubernurValues,
  dpmValues,
  dapilValues,
  calonDapilValues,
  zigZagMod,
} from "./common.js";

import gDaftarPemilih from "./daftarPemilih.js";
import gDaftarPilihan from "./daftarPilihan.js";

const daftarPemilihFile = "./test-data/daftar-pemilih.csv";
const daftarPilihanFile =
  "./test-data/PEMILIHAN UMUM MAHASISWA ELECTRONIC VOTE FAKULTAS TEKNIK UNIVERSITAS TRUNOJOYO MADURA .csv";

const generateBasedOnFixturesMatrix = (dapil) => {
  const daftarPemilih = gDaftarPemilih.generate(64, {
    nim: (i) => faker.random.alphaNumeric(12),
    dapil: (i) => dapil,
  });

  const userDapilList = [];
  for (let i = 0; i < 64; i++) {
    userDapilList.push(
      zigZagMod(i, 32)
        ? dapil
        : faker.random.arrayElement(
            dapilValues.filter((x) => x !== dapil)
          )
    );
  }

  const dapilList = [];
  for (let i = 0; i < 64; i++) {
    dapilList.push(
      zigZagMod(i, 64)
        ? dapil
        : faker.random.arrayElement(
            dapilValues.filter((x) => x !== dapil)
          )
    );
  }

  const chooseCalonDapilValue = (assignedDapil, i) => {
    if (zigZagMod(i, 64)) {
      if (dapil === assignedDapil) {
        return faker.random.arrayElement(calonDapilValues);
      } else {
        return "";
      }
    } else {
      if (dapilList[i] === assignedDapil) {
        return faker.random.arrayElement(calonDapilValues);
      } else {
        return "";
      }
    }
  };

  const daftarPilihan = gDaftarPilihan.generate(64, {
    email: (i) =>
      (zigZagMod(i, 2) ? daftarPemilih[i].NIM : faker.random.alphaNumeric(12)) +
      (zigZagMod(i, 4) ? "@student.trunojoyo.ac.id" : "@trunojoyo.ac.id"),
    waktu: (i) =>
      zigZagMod(i, 16)
        ? formatDate(faker.date.between(startDate, middleDate))
        : formatDate(faker.date.between(wrongStartDate, wrongEndDate)),
    dapil: (i) => userDapilList[i],
    persetujuan: (i) => "Setuju",
    gubernur: (i) => faker.random.arrayElement(gubernurValues),
    dpm: (i) => faker.random.arrayElement(dpmValues),
    himatif: (i) => chooseCalonDapilValue("TEKNIK INFORMATIKA", i),
    himatro: (i) => chooseCalonDapilValue("TEKNIK ELEKTRO", i),
    hmti: (i) => chooseCalonDapilValue("TEKNIK INDUSTRI", i),
  });

  return {
    daftarPemilih,
    daftarPilihan,
  };
};

const generateSah = (count) => {
  const daftarPemilih = gDaftarPemilih.generate(count, {
    nim: (i) => faker.random.alphaNumeric(12),
    dapil: (i) => faker.random.arrayElement(dapilValues),
  });

  const daftarPilihan = gDaftarPilihan.generate(daftarPemilih.length, {
    waktu: (i) => formatDate(faker.date.between(startDate, middleDate)),
    email: (i) => daftarPemilih[i].NIM + "@student.trunojoyo.ac.id", // + faker.internet.domainName(),
    persetujuan: (i) => "Setuju",
    dapil: (i) => daftarPemilih[i].PRODI,
    gubernur: (i) => faker.random.arrayElement(gubernurValues),
    dpm: (i) => faker.random.arrayElement(dpmValues),
    himatif: (i) =>
      daftarPemilih[i].PRODI === "TEKNIK INFORMATIKA"
        ? faker.random.arrayElement(calonDapilValues)
        : "",
    himatro: (i) =>
      daftarPemilih[i].PRODI == "TEKNIK ELEKTRO"
        ? faker.random.arrayElement(calonDapilValues)
        : "",
    hmti: (i) =>
      daftarPemilih[i].PRODI == "TEKNIK INDUSTRI"
        ? faker.random.arrayElement(calonDapilValues)
        : "",
  });

  return {
    daftarPemilih,
    daftarPilihan,
  };
};

let allDaftarPemilih = [];
let allDaftarPilihan = [];
for (let times = 0; times < 10; times++) {
  for (let i = 0; i < dapilValues.length; i++) {
    const o = generateBasedOnFixturesMatrix(dapilValues[i]);
    allDaftarPemilih.push(...o.daftarPemilih);
    allDaftarPilihan.push(...o.daftarPilihan);
  }
}

try {
  fs.writeFileSync(
    daftarPemilihFile,
    stringify(allDaftarPemilih, {
      columns: gDaftarPemilih.columns,
      header: true,
      quoted: true,
      quoted_empty: true,
    })
  );
} catch (err) {
  console.error(err);
}

try {
  fs.writeFileSync(
    daftarPilihanFile,
    stringify(allDaftarPilihan, {
      columns: gDaftarPilihan.columns,
      header: true,
      quoted: true,
      quoted_empty: true,
    })
  );
} catch (err) {
  console.error(err);
}
