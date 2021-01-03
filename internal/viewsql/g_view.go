package viewsql

// Target008Create contains SQL to create view vw_hasil_dpm
var Target008Create = `
CREATE VIEW "evote"."vw_hasil_dpm" AS

WITH

-- Penghitungan (aggregasi) suara sah
suara_sah AS (
  SELECT
    A.calon AS suara,
    TRUE AS sah,
    (
      SELECT
        COUNT(*)
      FROM evote.vw_pilihan B
      WHERE
        B.suara_non_dapil_sah = TRUE AND
        B.pilihan_dpm = A.calon
    ) AS jumlah
  FROM evote.ref_calon_dpm A
  ORDER BY A.calon
),

-- Data helper untuk alasan tidak sah
alasan AS (
  SELECT UNNEST(ARRAY[
    '(a) Tidak terdaftar dalam daftar pemilih',
    '(b) Di luar waktu pemungutan suara'
  ]) AS alasan
),

-- Penghitungan (aggregasi) suara tidak sah
suara_tidak_sah AS (
  SELECT
    A.alasan AS suara,
    FALSE AS sah,
    (
      SELECT
        COUNT(*)
      FROM evote.vw_pilihan B
      WHERE
        B.suara_non_dapil_sah = FALSE AND
        B.alasan_suara_non_dapil_tidak_sah = A.alasan
    ) AS jumlah
  FROM alasan A
),

-- Penggabungan view suara sah dan tidak sah
gabungan AS (
  SELECT * FROM suara_sah
  UNION ALL
  SELECT * FROM suara_tidak_sah
)

-- Penghitungan persentase
SELECT
  A.*,
  CASE
    WHEN A.sah THEN CAST((A.jumlah * 100) AS FLOAT) / NULLIF((SELECT COUNT(*) FROM evote.vw_pilihan WHERE suara_non_dapil_sah), 0)
    ELSE NULL
  END AS persentase_dari_suara_sah,
  CAST((A.jumlah * 100) AS FLOAT) / NULLIF((SELECT COUNT(*) FROM evote.vw_pilihan), 0) AS persentase_dari_seluruh_suara
FROM gabungan A`

// Target008Drop contains SQL to drop view vw_hasil_dpm
var Target008Drop = `DROP VIEW IF EXISTS "evote"."vw_hasil_dpm"`

// Target009Create contains SQL to create view vw_hasil_gubernur
var Target009Create = `
CREATE VIEW "evote"."vw_hasil_gubernur" AS

WITH

-- Penghitungan (aggregasi) suara sah
suara_sah AS (
  SELECT
    A.calon AS suara,
    TRUE AS sah,
    (
      SELECT
        COUNT(*)
      FROM evote.vw_pilihan B
      WHERE
        B.suara_non_dapil_sah = TRUE AND
        B.pilihan_gubernur = A.calon
    ) AS jumlah
  FROM evote.ref_calon_gubernur A
  ORDER BY A.calon
),

-- Data helper untuk alasan tidak sah
alasan AS (
  SELECT UNNEST(ARRAY[
    '(a) Tidak terdaftar dalam daftar pemilih',
    '(b) Di luar waktu pemungutan suara'
  ]) AS alasan
),

-- Penghitungan (aggregasi) suara tidak sah
suara_tidak_sah AS (
  SELECT
    A.alasan AS suara,
    FALSE AS sah,
    (
      SELECT
        COUNT(*)
      FROM evote.vw_pilihan B
      WHERE
        B.suara_non_dapil_sah = FALSE AND
        B.alasan_suara_non_dapil_tidak_sah = A.alasan
    ) AS jumlah
  FROM alasan A
),

-- Penggabungan view suara sah dan tidak sah
gabungan AS (
  SELECT * FROM suara_sah
  UNION ALL
  SELECT * FROM suara_tidak_sah
)

-- Penghitungan persentase
SELECT
  A.*,
  CASE
    WHEN A.sah THEN CAST((A.jumlah * 100) AS FLOAT) / NULLIF((SELECT COUNT(*) FROM evote.vw_pilihan WHERE suara_non_dapil_sah), 0)
    ELSE NULL
  END AS persentase_dari_suara_sah,
  CAST((A.jumlah * 100) AS FLOAT) / NULLIF((SELECT COUNT(*) FROM evote.vw_pilihan), 0) AS persentase_dari_seluruh_suara
FROM gabungan A`

// Target009Drop contains SQL to drop view vw_hasil_gubernur
var Target009Drop = `DROP VIEW IF EXISTS "evote"."vw_hasil_gubernur"`

// Target010Create contains SQL to create view vw_hasil_hmp
var Target010Create = `
CREATE VIEW "evote"."vw_hasil_hmp" AS

WITH

-- Data helper untuk daftar calon setiap dapil (dikumpulkan dalam satu
-- view)
all_calon AS (
  --SELECT 'TEKNIK INFORMATIKA' AS dapil, calon FROM evote.ref_calon_himatif
  --UNION ALL
  SELECT 'TEKNIK ELEKTRO' AS dapil, calon FROM evote.ref_calon_himatro
  UNION ALL
  SELECT 'TEKNIK INDUSTRI' AS dapil, calon FROM evote.ref_calon_hmti
),

-- Menggabungkan kolom-kolom pilihan tiap dapil menjadi satu, dengan cara menampilkan
-- hanya kolom pilihan dapil yang bersesuaian dengan dapil pemilih.
dapil_normalized AS (
  SELECT
    A.nim_email_hashed,
    A.suara_dapil_sah,
    A.alasan_suara_dapil_tidak_sah,
    A.dapil_tercatat AS dapil,
    CASE
      --WHEN A.dapil_tercatat = 'TEKNIK INFORMATIKA' THEN A.pilihan_himatif
      WHEN A.dapil_tercatat = 'TEKNIK ELEKTRO' THEN A.pilihan_himatro
      WHEN A.dapil_tercatat = 'TEKNIK INDUSTRI' THEN A.pilihan_hmti
      ELSE NULL
    END AS pilihan_hmp
  FROM evote.vw_pilihan A
),

-- Penghitungan (aggregasi) suara sah untuk setiap dapil
suara_sah AS (
  SELECT
    A.dapil,
    A.calon AS suara,
    TRUE AS sah,
    (
      SELECT
        COUNT(*)
      FROM dapil_normalized B
      WHERE
        B.dapil = A.dapil AND
        B.suara_dapil_sah = TRUE AND
        B.pilihan_hmp = A.calon
    ) AS jumlah
  FROM all_calon A
  ORDER BY A.dapil, A.calon
),

-- Data helper untuk alasan tidak sah
alasan AS (
  SELECT UNNEST(ARRAY[
    -- Tidak ikut ditampilkan di sini, karena bila tidak ada dalam daftar
    -- pemilih maka tidak dapat diverifikasi pemilih tersebut berasal dari
    -- dapil mana
    -- '(a) Tidak terdaftar dalam daftar pemilih',
    '(b) Dapil tidak sesuai',
    '(c) Di luar waktu pemungutan suara'
  ]) AS alasan
),

-- Data helper untuk data dapil
fak AS (
  SELECT UNNEST(ARRAY[
    --'TEKNIK INFORMATIKA',
    'TEKNIK ELEKTRO',
    'TEKNIK INDUSTRI'
    -- 'SISTEM INFORMASI',
    -- 'TEKNIK MESIN',
    -- 'TEKNIK MEKATRONIKA'
  ]) AS dapil
),

-- Kombinasi data dapil dan alasan tidak sah
alasan_dapil AS (
  SELECT A.dapil, B.alasan FROM fak A, alasan B
),

-- Penghitungan (aggregasi) suara tidak sah untuk setiap dapil
suara_tidak_sah AS (
  SELECT
    A.dapil,
    A.alasan AS suara,
    FALSE AS sah,
    (
      SELECT
        COUNT(*)
      FROM dapil_normalized B
      WHERE
        B.dapil = A.dapil AND
        B.suara_dapil_sah = FALSE AND
        B.alasan_suara_dapil_tidak_sah = A.alasan
    ) AS jumlah
  FROM alasan_dapil A
),

-- Penggabungan view suara sah dan tidak sah
gabungan AS (
  SELECT * FROM suara_sah
  UNION ALL
  SELECT * FROM suara_tidak_sah
)

-- Penghitungan persentase
(SELECT
  A.*,
  CASE
    WHEN A.sah THEN CAST((A.jumlah * 100) AS FLOAT) / NULLIF((SELECT COUNT(*) FROM evote.vw_pilihan B WHERE B.dapil_tercatat = A.dapil AND B.suara_dapil_sah), 0)
  END AS persentase_dari_suara_sah_dalam_dapil,
  CAST((A.jumlah * 100) AS FLOAT) / NULLIF((SELECT COUNT(*) FROM evote.vw_pilihan B WHERE B.dapil_tercatat = A.dapil), 0) AS persentase_dari_seluruh_suara_dalam_dapil
FROM gabungan A
ORDER BY A.dapil, A.sah DESC, A.suara)

UNION ALL

(SELECT
  NULL AS dapil,
  '(a) Tidak terdaftar dalam daftar pemilih' AS suara,
  FALSE AS sah,
  (
    SELECT COUNT(*)
    FROM dapil_normalized
    WHERE alasan_suara_dapil_tidak_sah = '(a) Tidak terdaftar dalam daftar pemilih'
  ) AS jumlah,
  NULL AS persentase_dari_suara_sah_dalam_dapil,
  NULL AS persentase_dari_seluruh_suara_dalam_dapil)`

// Target010Drop contains SQL to drop view vw_hasil_hmp
var Target010Drop = `DROP VIEW IF EXISTS "evote"."vw_hasil_hmp"`

// Target011Create contains SQL to create view vw_pilihan
var Target011Create = `
CREATE VIEW "evote"."vw_pilihan" AS

WITH convert_to_timestamp AS (
  SELECT
    nim_email_hashed,
    -- Konversi tipe data waktu dari string ke timestamp
    TO_TIMESTAMP(waktu_str, 'YYYY/MM/DD HH:MI:SS AM') AS waktu,
    jenis_email,
    dapil,
    pilihan_gubernur,
    pilihan_dpm,
    --pilihan_himatif,
    pilihan_himatro,
    pilihan_hmti
  FROM evote.daftar_pilihan
),

latest_data AS (
  SELECT
    nim_email_hashed,
    MAX(waktu) AS waktu
  FROM convert_to_timestamp
  GROUP BY nim_email_hashed
),

validasi AS (
  SELECT
    A.waktu,

    -- dalam_waktu_pemungutan akan bernilai TRUE apabila waktu pengisian form
    -- oleh pemilih berada dalam batas waktu yang ditentukan
    A.waktu BETWEEN
      TO_TIMESTAMP('2021/01/04 8:00:00 AM', 'YYYY/MM/DD HH:MI:SS AM') AND
      TO_TIMESTAMP('2021/01/04 1:00:00 PM', 'YYYY/MM/DD HH:MI:SS AM')
    AS dalam_waktu_pemungutan,

    A.nim_email_hashed,
    B.jenis_email,
    B.dapil AS input_dapil,
    C.dapil AS dapil_tercatat,

    -- terdaftar akan bernilai TRUE apabila email/NIM yang bersangkutan
    -- terdaftar dalam daftar pemilih dan email yang digunakan memiliki domain
    -- student.trunojoyo.ac.id
    (C.nim_email_hashed IS NOT NULL AND B.jenis_email = 'student') AS terdaftar,

    -- dapil_benar akan bernilai TRUE apabila data dapil yang diinputkan
    -- pemilih pada Google Forms sesuai dengan yang ada pada data
    -- daftar pemilih.
    B.dapil = C.dapil AS dapil_benar,
    
    -- dapil_sesuai akan bernilai TRUE apabila pemilih memberikan pilihannya
    -- pada dapil yang sesuai dengan fakultas pemilih yang tercatat di daftar
    -- pemilih.
    --
    -- Kasus berikut ini KEMUNGKINAN KECIL terjadi karena sudah dibatasi pada
    -- Google Forms.
    -- Apabila ternyata pemilih selain memilih pada dapil fakultas juga
    -- memilih pada dapil lainnya, maka pilihan pada dapil lainnya tersebut
    -- akan diabaikan.
    CASE
      --WHEN C.dapil = 'TEKNIK INFORMATIKA' THEN B.pilihan_himatif IS NOT NULL
      WHEN C.dapil = 'TEKNIK ELEKTRO' THEN B.pilihan_himatro IS NOT NULL
      WHEN C.dapil = 'TEKNIK INDUSTRI' THEN B.pilihan_hmti IS NOT NULL
    END AS dapil_sesuai,
    
    B.pilihan_gubernur,
    B.pilihan_dpm,
    --B.pilihan_himatif,
    B.pilihan_himatro,
    B.pilihan_hmti
  FROM latest_data A
  LEFT JOIN convert_to_timestamp B ON (A.nim_email_hashed = B.nim_email_hashed AND A.waktu = B.waktu)
  LEFT JOIN evote.daftar_pemilih C ON (A.nim_email_hashed = C.nim_email_hashed)
)

-- Implementasi penentuan suara sah dan tidak sah (beserta alasan tidak
-- sahnya)
SELECT
  CASE
      WHEN A.terdaftar = FALSE THEN FALSE
      WHEN A.dalam_waktu_pemungutan = FALSE THEN FALSE
      ELSE TRUE
  END AS suara_non_dapil_sah,

  CASE
      WHEN A.terdaftar = FALSE THEN '(a) Tidak terdaftar dalam daftar pemilih'
      WHEN A.dalam_waktu_pemungutan = FALSE THEN '(b) Di luar waktu pemungutan suara'
      ELSE NULL
  END AS alasan_suara_non_dapil_tidak_sah,

  CASE
      WHEN A.terdaftar = FALSE THEN FALSE
      WHEN A.dapil_sesuai = FALSE THEN FALSE
      WHEN A.dalam_waktu_pemungutan = FALSE THEN FALSE
      ELSE TRUE
  END AS suara_dapil_sah,

  CASE
      WHEN A.terdaftar = FALSE THEN '(a) Tidak terdaftar dalam daftar pemilih'
      WHEN A.dapil_sesuai = FALSE THEN '(b) Dapil tidak sesuai'
      WHEN A.dalam_waktu_pemungutan = FALSE THEN '(c) Di luar waktu pemungutan suara'
      ELSE NULL
  END AS alasan_suara_dapil_tidak_sah,
  A.*
FROM validasi A`

// Target011Drop contains SQL to drop view vw_pilihan
var Target011Drop = `DROP VIEW IF EXISTS "evote"."vw_pilihan"`
