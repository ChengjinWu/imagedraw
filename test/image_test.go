package test

import (
	"fmt"
	"github.com/BurntSushi/graphics-go/graphics"
	"github.com/disintegration/imaging"
	"image"
	"image/draw"
	"imagedraw"
	"testing"
)

func GetHeadImageRGBA(iamgePath string) (*image.RGBA, error) {
	img, err := imaging.Open(iamgePath)
	if err != nil {
		return nil, err
	}
	if img.Bounds().Dx() > img.Bounds().Dy() {
		img = imaging.CropCenter(img, img.Bounds().Dy(), img.Bounds().Dy())
	} else {
		img = imaging.CropCenter(img, img.Bounds().Dx(), img.Bounds().Dx())
	}
	imgRGBA := image.NewRGBA(image.Rect(0, 0, 190, 190))
	err = graphics.Scale(imgRGBA, img)
	if err != nil {
		return nil, err
	}
	return imgRGBA, nil
}
func Test_Image(T *testing.T) {

	textBrush, err := imagedraw.NewTextBrush("华文仿宋 加粗.TTF", 26, image.White, 1000)
	if err != nil {
		fmt.Println(err)
	}
	backgroundImg, err := imaging.Open("background.png")
	if err != nil {
		fmt.Println(err)
	}
	imgRGBA, err := GetHeadImageRGBA("1.jpg")
	if err != nil {
		fmt.Println(err)
	}
	img2RGBA, err := GetHeadImageRGBA("2.jpg")
	if err != nil {
		fmt.Println(err)
	}
	m := image.NewRGBA(backgroundImg.Bounds())
	var x0, y0 int
	draw.Draw(m, backgroundImg.Bounds(), backgroundImg, image.ZP, draw.Src)
	x0, y0 = 89, 360
	draw.DrawMask(m, image.Rect(x0, y0, x0+200, y0+200), imgRGBA, image.ZP, &imagedraw.CircleMask{image.Pt(95, 95), 95}, image.ZP, draw.Over)
	x0, y0 = 361, 342
	draw.DrawMask(m, image.Rect(x0, y0, x0+200, y0+200), img2RGBA, image.ZP, &imagedraw.CircleMask{image.Pt(95, 95), 95}, image.ZP, draw.Over)
	textBrush.FontSize = 54
	//textBrush.DrawFontOnRGBA(m, image.Pt(100,100),1000,"习近平指出，总统先生不久前对中国进行了十分成功的国事访问，我们就中美关系和共同关心的重大问题深入交换意见，达成多方面重要共识，对推动中美关系保持健康稳定发展具有重要意义。")
	textBrush.FontSize = 26
	textBrush.DrawFontOnRGBA(m, image.Pt(20, 20), `尊敬的 {{.To}}:
					这里是您的邀请函,你可以凭借此文件来我公司进行商谈你可以凭借此文件来我公司进行商谈你可以凭借此文件来我公司进行商谈你可以凭借此文件来我公司进行商谈你可以凭借此文件来我公司进行商谈你可以凭借此文件来我公司进行商谈你可以凭借此文件来我公司进行商谈你可以凭借此文件来我公司进行商谈.



                重要: {{.From}}



                     {{.Date}}`)

	imaging.Save(m, "dst.jpg")
}
