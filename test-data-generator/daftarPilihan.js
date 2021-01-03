const columns = [
  "Timestamp",
  "Username",
  "SAYA MENGGUNAKAN HAK PILIH DALAM PEMILU E-VOTE SECARA SADAR DAN TANPA TEKANAN DARI SIAPAPUN:",
  "PROGRAM STUDI :",
  "Silahkan pilih calon Gubernur dan Wakil Gubernur FT UTM",
  "Silahkan pilih calon Anggota DPM FT UTM",
  //"Silahkan pilih calon KETUA UMUM dan WAKIL KETUA UMUM HIMATIF FT UTM ",
  "Silahkan pilih calon KETUA UMUM dan WAKIL KETUA UMUM HIMATRO FT UTM ",
  "Silahkan pilih calon KETUA UMUM dan WAKIL KETUA UMUM HMTI FT UTM",
];

const generate = (
  count,
  { waktu, email, persetujuan, dapil, gubernur, dpm, himatro, hmti }
) => {
  const records = [];
  for (let i = 0; i < count; i++) {
    records.push({
      "Timestamp": waktu(i),
      Username: email(i),
      "SAYA MENGGUNAKAN HAK PILIH DALAM PEMILU E-VOTE SECARA SADAR DAN TANPA TEKANAN DARI SIAPAPUN:": persetujuan(
        i
      ),
      "PROGRAM STUDI :": dapil(i),
      "Silahkan pilih calon Gubernur dan Wakil Gubernur FT UTM": gubernur(i),
      "Silahkan pilih calon Anggota DPM FT UTM": dpm(i),
      //"Silahkan pilih calon KETUA UMUM dan WAKIL KETUA UMUM HIMATIF FT UTM ": himatif(
        //i
      //),
      "Silahkan pilih calon KETUA UMUM dan WAKIL KETUA UMUM HIMATRO FT UTM ": himatro(
        i
      ),
      "Silahkan pilih calon KETUA UMUM dan WAKIL KETUA UMUM HMTI FT UTM": hmti(
        i
      ),
    });
  }
  return records;
};

export default {
  columns,
  generate,
};
