package main

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
)

type Profile struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Image    string `json:"image"`
}

func main() {
	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/static/", fs)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var profiles = []Profile{
		{
			Username: "reposter",
			Image:    "https://imgur.com/fQDQJeg.jpg",
			Avatar:   "https://i.imgur.com/YgJ4qks.jpg",
		},
	}

	for _, profile := range profiles {
		img := Run(profile.Username, profile.Image, profile.Avatar)

		buf := new(bytes.Buffer)
		err := png.Encode(buf, img)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "image/jpeg")
		w.Write(buf.Bytes())

		return
	}
}

// dimensions
const AvatarSize = 48
const BoxBorder = 10
const AvatarPositionX = 10
const UsernamePositionX = AvatarPositionX + AvatarSize + 5
const FontPoints = 35

func BoxHeight(h float64) float64 {
	return float64((h * 6) / 100)
}

func CenterAvatar(imageHeight int) float64 {
	return (float64(imageHeight) - AvatarSize) - ((BoxHeight(float64(imageHeight)) - AvatarSize) / 2)
}

func BoxWidth(dc *gg.Context, username string) float64 {
	fontWidth, _ := dc.MeasureString(username)
	return fontWidth + AvatarPositionX + UsernamePositionX + 10
}
func CenterUserNameY(imageHeight int, fontSize float64) float64 {
	return CenterAvatar(imageHeight) + (fontSize / 2)
}

func Run(username, image, avatar string) image.Image {
	// img
	img := getImage(image)
	g := img.Bounds()
	imageWidth := g.Dx()
	imageHeight := g.Dy()

	dc := gg.NewContext(imageWidth, imageHeight)
	_, fontHeight := dc.MeasureString(username)
	loadFont(dc)
	dc.DrawImage(img, 0, 0)
	drawBox(dc, float64(imageHeight), username)

	drawAvatar(dc, avatar, AvatarPositionX, CenterAvatar(imageHeight))
	drawUsername(dc, username, float64(UsernamePositionX), CenterUserNameY(imageHeight, fontHeight))

	return dc.Image()
}

func getImage(imageURL string) image.Image {
	// get image
	resp, _ := http.Get(imageURL)
	defer resp.Body.Close()

	// decode img
	img, _ := jpeg.Decode(resp.Body)

	return img
}

func drawBox(dc *gg.Context, imageHeight float64, username string) {
	boxHeight := BoxHeight(imageHeight)
	boxWidth := BoxWidth(dc, username)

	borderRightSpacing := 10.0

	//TODO: 350 magic number
	dc.SetRGB255(255, 255, 255)
	x := 0.0
	y := imageHeight - boxHeight

	dc.DrawRectangle(x, y, float64(boxWidth-borderRightSpacing), float64(boxHeight))
	dc.FillPreserve()
	dc.DrawRoundedRectangle(x, y, float64(boxWidth), float64(boxHeight), BoxBorder)
	dc.FillPreserve()
	dc.Clip()
}

func drawAvatar(dc *gg.Context, avatarURL string, x, y float64) {
	// get avatar
	avatarResp, _ := http.Get(avatarURL)
	defer avatarResp.Body.Close()
	avatarBody, _ := ioutil.ReadAll(avatarResp.Body)

	// decode avatar
	avatar, _, _ := image.Decode(bytes.NewReader(avatarBody))

	dstImageFill := imaging.Fill(avatar, AvatarSize, AvatarSize, imaging.Center, imaging.Lanczos)
	g := dstImageFill.Bounds()
	avatarWidth := g.Dx()
	avatarHeight := g.Dy()

	dc.DrawRoundedRectangle(x, y, float64(avatarWidth), float64(avatarHeight), float64(avatarHeight)/2)
	dc.Clip()
	dc.DrawImage(dstImageFill, int(x), int(y))
	dc.ResetClip()
}

func drawUsername(dc *gg.Context, username string, x, y float64) {
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored(username, x, y, 0, 1)
}

func loadFont(dc *gg.Context) {
	if err := dc.LoadFontFace("./static_Oswald.ttf", FontPoints); err != nil {
		panic(err)
	}
}
