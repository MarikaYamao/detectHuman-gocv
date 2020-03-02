package main
import (
    "encoding/json"
    "fmt"
    "image"
    "image/color"
    "gocv.io/x/gocv"
)

func main() {
    // Detected by passing learned data to classifier
    // The data used this time is the data originally included in opencv / data
    harrcascade := "haarcascade_frontalface_default.xml"
    classifier := gocv.NewCascadeClassifier()
    classifier.Load(harrcascade)
    defer classifier.Close()

    // color for the rect when faces detected
    blue := color.RGBA{0, 0, 255, 0}

    // read image
    img := gocv.IMRead("images/person_01.jpg", 0)

    //resize image
    fact := float64(400) / float64(img.Cols())
    newY := float64(img.Rows()) * fact
    gocv.Resize(img, &img, image.Point{X: 400, Y: int(newY)}, 0, 0, 1)

    // detect people in image
    rects := classifier.DetectMultiScale(img)

    // print found points
    printStruct(rects)

    // draw a rectangle around each face on the original image,
    // along with text identifing as "Human"
    for _, r := range rects {
        gocv.Rectangle(&img, r, blue, 2)
        size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
        pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
        gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1, blue, 2)
    }
    if ok := gocv.IMWrite("gocv_detected.jpg", img); !ok {
        fmt.Println("Error")
    }
}
func printStruct(i interface{}) {
  b, err := json.Marshal(i)
  if err != nil {
      fmt.Println(err)
      return
  }
  fmt.Println(string(b))
}