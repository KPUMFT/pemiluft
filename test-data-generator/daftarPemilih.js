const columns = ["NIM", "PRODI"];
const generate = (count, { nim, dapil }) => {
  const records = [];
  for (let i = 0; i < count; i++) {
    records.push({
      NIM: nim(i),
      PRODI: dapil(i),
    });
  }
  return records;
};

export default {
  columns,
  generate,
};
