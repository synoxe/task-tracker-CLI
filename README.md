🚀 Task Tracker CLI






💻 Terminalden hızlı ve sade görev yönetimi için geliştirilmiş bir CLI uygulaması.
Minimal yapısı sayesinde odak kaybetmeden üretken kalmanı sağlar.

✨ Özellikler
⚡ Hızlı CLI deneyimi
➕ Görev ekleme
✏️ Güncelleme
❌ Silme
🔄 Durum yönetimi
📃 Listeleme ve filtreleme
💾 JSON ile kalıcı veri saklama
📸 Demo
$ ./task-cli add "Go öğren"
✔ Görev eklendi (ID: 1)

$ ./task-cli list
[1] Go öğren (todo)

$ ./task-cli mark-done 1
✔ Görev tamamlandı

$ ./task-cli list done
[1] Go öğren (done)
🛠️ Kurulum
git clone https://github.com/synoxe/task-tracker-CLI.git
cd task-tracker-CLI
go build -o task-cli

Alternatif:

go run main.go
⚙️ Kullanım
➕ Görev ekle
./task-cli add "Yeni görev"
✏️ Güncelle
./task-cli update <id> "Yeni açıklama"
❌ Sil
./task-cli delete <id>
🔄 Durum değiştir
./task-cli mark-in-progress <id>
./task-cli mark-done <id>
📃 Listele
./task-cli list
./task-cli list todo
./task-cli list in-progress
./task-cli list done
📁 Veri Yapısı

Tüm görevler tasks.json içinde tutulur:

[
  {
    "id": 1,
    "description": "Go öğren",
    "status": "todo",
    "createdAt": "2026-04-17T12:00:00Z",
    "updatedAt": null
  }
]
🧠 Mantık
Her işlem JSON dosyasını günceller
ID’ler otomatik artar
Durum bazlı filtreleme yapılır
Basit ama genişletilebilir yapı
📦 Teknolojiler
🟦 Go (Golang)
📄 JSON
🖥️ CLI
🗺️ Roadmap
 Renkli terminal çıktısı
 Config dosyası desteği
 Export (CSV / Markdown)
 Tarihe göre filtreleme
 TUI (terminal UI) versiyonu
🤝 Katkı

Katkı yapmak istersen:

fork → branch → commit → push → PR
📄 Lisans

MIT License

⭐ Destek

Projeyi beğendiysen ⭐ bırakmayı unutma!

https://roadmap.sh/projects/task-tracker
