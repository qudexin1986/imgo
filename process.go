package imgo

import (
	"errors"
)

//input a image matrix as src , return a image matrix by sunseteffect process
func SunsetEffect(src [][][]uint8)(imgMatrix [][][]uint8 , err error) {
	imgMatrix = src
	
	height:=len(imgMatrix)
	width:=len(imgMatrix[0])
	if height == 0 || width == 0 {
		err = errors.New("The input of matrix is illegal!")
	}
	
	for i:=0;i<height;i++{
		for j:=0;j<width;j++{
			imgMatrix[i][j][1] = uint8( float64(imgMatrix[i][j][1]) * 0.7 )
			imgMatrix[i][j][2] = uint8( float64(imgMatrix[i][j][2]) * 0.7 )
		}
	}
	
	return
}

// input a image as src , return a image matrix by negativefilmeffect process
func NegativeFilmEffect(src [][][]uint8)(imgMatrix [][][]uint8 , err error) {
	imgMatrix = src
	
	height:=len(imgMatrix)
	width:=len(imgMatrix[0])
	if height == 0 || width == 0 {
		err = errors.New("The input of matrix is illegal!")
	}
	
	for i:=0;i<height;i++{
		for j:=0;j<width;j++{
			imgMatrix[i][j][0] = 255 - imgMatrix[i][j][0]
			imgMatrix[i][j][1] = 255 - imgMatrix[i][j][1]
			imgMatrix[i][j][2] = 255 - imgMatrix[i][j][2]
		}
	}
	
	return
}

func AdjustBrightness(src [][][]uint8 , light float64)(imgMatrix [][][]uint8 , err error) {
	imgMatrix = src
	
	if light <= 0{
		err = errors.New("value of light must be more than 0")
		return
	}
	
	height:=len(imgMatrix)
	width:=len(imgMatrix[0])
	if height == 0 || width == 0 {
		err = errors.New("The input of matrix is illegal!")
	}
	
	for i:=0;i<height;i++{
		for j:=0;j<width;j++{
			imgMatrix[i][j][0] = uint8(float64(imgMatrix[i][j][0])*light)
			imgMatrix[i][j][1] = uint8(float64(imgMatrix[i][j][1])*light)
			imgMatrix[i][j][2] = uint8(float64(imgMatrix[i][j][2])*light)
		}
	}
	
	return
}

// fuse two images(filepath) and the size of new image is as src1
func ImageFusion(src1 string , src2 string)(imgMatrix [][][]uint8 , err error) {
	imgMatrix1,err1 := Read(src1)
	
	if err1 != nil {
		err = err1
		return 
	}
	
	
	height:=len(imgMatrix1)
	width:=len(imgMatrix1[0])
	
	imgMatrix2,err2 := ResizeForMatrix(src2,width,height)
	
	if err2 != nil {
		err = err2
		return
	}
		
	for i:=0;i<height;i++{
		for j:=0;j<width;j++{
			imgMatrix1[i][j][0] = uint8(float64(imgMatrix1[i][j][0])*0.5)+uint8(float64(imgMatrix2[i][j][0])*0.5)
			imgMatrix1[i][j][1] = uint8(float64(imgMatrix1[i][j][1])*0.5)+uint8(float64(imgMatrix2[i][j][1])*0.5)
			imgMatrix1[i][j][2] = uint8(float64(imgMatrix1[i][j][2])*0.5)+uint8(float64(imgMatrix1[i][j][2])*0.5)
		}
	}
	imgMatrix = imgMatrix1
	return	
}


func VerticalMirror(src [][][]uint8)(imgMatrix [][][]uint8 , err error){
	height:=len(src)
	width:=len(src[0])
	if height == 0 || width == 0 {
		err = errors.New("The input of matrix is illegal!")
	}
	
	newwidth:=width*2
	imgMatrix=NewRGBAMatrix(height,newwidth)
	
	for i:=0;i<height;i++{
		for j:=0;j<width;j++{
		imgMatrix[i][j][0] = src[i][j][0]
		imgMatrix[i][j][1] = src[i][j][1]
		imgMatrix[i][j][2] = src[i][j][2]
		imgMatrix[i][j][3] = src[i][j][3]
		}
	}
	
	
	for i:=0;i<height;i++{
		for j:=width;j<newwidth;j++{
			imgMatrix[i][j][0] = imgMatrix[i][newwidth-j-1][0]
			imgMatrix[i][j][1] = imgMatrix[i][newwidth-j-1][1]
			imgMatrix[i][j][2] = imgMatrix[i][newwidth-j-1][2]
			imgMatrix[i][j][3] = imgMatrix[i][newwidth-j-1][3]
		}
	}
	
	return
}

func HorizontalMirror(src [][][]uint8)(imgMatrix [][][]uint8 , err error){
	height:=len(src)
	width:=len(src[0])
	if height == 0 || width == 0 {
		err = errors.New("The input of matrix is illegal!")
	}
	
	newheight:=height*2
	imgMatrix=NewRGBAMatrix(newheight,width)
	
	for i:=0;i<height;i++{
		for j:=0;j<width;j++{
		imgMatrix[i][j][0] = src[i][j][0]
		imgMatrix[i][j][1] = src[i][j][1]
		imgMatrix[i][j][2] = src[i][j][2]
		imgMatrix[i][j][3] = src[i][j][3]
		}
	}
	
	
	for i:=height;i<newheight;i++{
		for j:=0;j<width;j++{
			imgMatrix[i][j][0] = imgMatrix[newheight-i-1][j][0]
			imgMatrix[i][j][1] = imgMatrix[newheight-i-1][j][1]
			imgMatrix[i][j][2] = imgMatrix[newheight-i-1][j][2]
			imgMatrix[i][j][3] = imgMatrix[newheight-i-1][j][3]
		}
	}
	
	return
}


func VerticalMirrorPart(src [][][]uint8)(imgMatrix [][][]uint8 , err error){
	imgMatrix = src
	
	height:=len(imgMatrix)
	width:=len(imgMatrix[0])
	if height == 0 || width == 0 {
		err = errors.New("The input of matrix is illegal!")
	}
	
	mirror_w:=width/2
	
	for i:=0;i<height;i++{
		for j:=0;j<mirror_w;j++{
			imgMatrix[i][j][0] = imgMatrix[i][width-j-1][0]
			imgMatrix[i][j][1] = imgMatrix[i][width-j-1][1]
			imgMatrix[i][j][2] = imgMatrix[i][width-j-1][2]
		}
	}
	
	return
}


func HorizontalMirrorPart(src [][][]uint8)(imgMatrix [][][]uint8 , err error){
	imgMatrix = src
	
	height:=len(imgMatrix)
	width:=len(imgMatrix[0])
	if height == 0 || width == 0 {
		err = errors.New("The input of matrix is illegal!")
	}
	
	mirror_h:=height/2
	
	for i:=0;i<mirror_h;i++{
		for j:=0;j<width;j++{
		imgMatrix[height-i-1][j][0] = imgMatrix[i][j][0]
		imgMatrix[height-i-1][j][1] = imgMatrix[i][j][1]
		imgMatrix[height-i-1][j][2] = imgMatrix[i][j][2]
		}
	}
	
	return
}


func RGB2Gray(src [][][]uint8)(imgMatrix [][][]uint8 , err error){
	imgMatrix = src
	
	height:=len(imgMatrix)
	width:=len(imgMatrix[0])
	if height == 0 || width == 0 {
		err = errors.New("The input of matrix is illegal!")
	}
	
	for i:=0;i<height;i++{
		for j:=0;j<width;j++{
		avg:=(imgMatrix[i][j][0]+imgMatrix[i][j][1]+imgMatrix[i][j][3])/3
		imgMatrix[i][j][0] = avg
		imgMatrix[i][j][1] = avg
		imgMatrix[i][j][2] = avg
		}
	}
	
	return
}