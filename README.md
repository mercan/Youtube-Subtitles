# YouTube Subtitles

**YouTube'da**, altyazısı olan veya otomatik olarak oluşturulmuş videolarda, istenilen kelime veya cümle hangi saat, 
dakika ve saniye aralığında geçtiğini göstererek, kullanıcının aradığı metni hızlıca bulmasını sağlar.


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

#### Altyazı Ara

```http
  GET /subtitles?url=YoutubeVideoURL&text=AranacakKelimeVeyaCümle
```

| Parameter |   Tip    |                        Açıklama                        |
|:---------:|:--------:|:------------------------------------------------------:|
|   `url`   | `string` |            **Gerekli**. Youtube Video URL.             |
|  `text`   | `string` | **Gerekli**. Video'da bulmak istediğiniz kelime/cümle. |
