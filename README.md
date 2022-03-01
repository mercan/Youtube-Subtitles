# YouTube Subtitles

**YouTube**'da altyazısı olan veya otomatik olarak oluşturulan altyazı olan videoda istediğiniz kelimenin hangi saat, dakika ve saniye de geçtiğini size gösterip aradığınız şeyi hızlıca bulmanızı sağlar.

Örnek [Request](http://localhost:3000/subtitles?url=https://www.youtube.com/watch?v=ARNNNmhSPME&text=değil)

## Bilgisayarınızda Çalıştırın

Projeyi klonlayın

```bash
  git clone https://github.com/mercan/Go-Youtube-Subtitles.git
```

Proje dizinine gidin

```bash
  cd Go-Youtube-Subtitles
```

Docker Build alın

```bash
    docker build -t go-youtube-subtitles .
```

Docker Image'ını çalıştırın

```bash
    docker run -p 3000:3000 -it go-youtube-subtitles
```
## API Kullanımı

#### Altyazı ara

```http
  GET /subtitles?{url}&{text}
```

| Parameter |   Tip    |                        Açıklama                        |
|:---------:|:--------:|:------------------------------------------------------:|
|   `url`   | `string` |            **Gerekli**. Youtube Video URL.             |
|  `text`   | `string` | **Gerekli**. Video'da bulmak istediğiniz kelime/cümle. |

## Kullanılan Teknolojiler

**Sunucu:** Go, Echo, Youtube-dl, Docker

  