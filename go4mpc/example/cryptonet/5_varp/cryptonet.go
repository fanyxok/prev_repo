package main

const ImageWidth = 28
const Stride = 2
const PaddedWidth = 30
const WindowWidth = 5
const OutputChannels = 5
const FcWidth = 100
const ConvWidth = 13
const ConvSize = 169
const FinalOutputChannels = 10

func main() {
	var image []sint8 = i8n(0, 784)
	var kernel []sint8 = i8n(0, 125)
	var pool []sint8 = i8n(0, 84500)
	var fc []sint8 = i8n(0, 1000)
	var conv_in []sint8 = make([]sint8, 900)
	conv_in[0] = 0
	conv_in[30] = 0
	conv_in[0] = 0
	conv_in[1] = 0
	for i := 1; i < 29; i = i + 1 {
		conv_in[i] = 0
		conv_in[i+30] = 0
		conv_in[30*i] = 0
		conv_in[30*i+1] = 0
	}
	conv_in[29] = 0
	conv_in[59] = 0
	conv_in[870] = 0
	conv_in[871] = 0
	conv_in[62] = image[0]
	for x := 1; x < 27; x = x + 1 {
		var img_pixel_5 int = 0 + x
		var in_pixel_5 int = 60 + x + 2
		conv_in[in_pixel_5] = image[img_pixel_5]
	}
	conv_in[89] = image[27]
	for y := 1; y < 27; y = y + 1 {
		var img_pixel_3 int = y*28 + 0
		var in_pixel_3 int = (y+2)*30 + 0 + 2
		conv_in[in_pixel_3] = image[img_pixel_3]
		for x := 1; x < 27; x = x + 1 {
			var img_pixel int = y*28 + x
			var in_pixel int = (y+2)*30 + x + 2
			conv_in[in_pixel] = image[img_pixel]
		}
		var img_pixel_4 int = y*28 + 27
		var in_pixel_4 int = (y+2)*30 + 27 + 2
		conv_in[in_pixel_4] = image[img_pixel_4]
	}
	conv_in[872] = image[756]
	for x := 1; x < 27; x = x + 1 {
		var img_pixel_6 int = 756 + x
		var in_pixel_6 int = 870 + x + 2
		conv_in[in_pixel_6] = image[img_pixel_6]
	}
	conv_in[899] = image[783]
	var output_conv_29 []sint8 = make([]sint8, 845)
	var kernel_19_conv_29 []sint8 = kernel[0:25]
	var out_conv_naive_38 []sint8 = make([]sint8, 169)
	var tmp_25_27_conv_naive_38 sint8 = 0
	var computed_21_23_25_27_conv_naive_38 sint8 = kernel_19_conv_29[0] * conv_in[0]
	tmp_25_27_conv_naive_38 = tmp_25_27_conv_naive_38 + computed_21_23_25_27_conv_naive_38
	for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
		var convPos_23_25_27_conv_naive_38 int = wx_conv_naive_38 + 0
		var computed_23_25_27_conv_naive_38 sint8 = kernel_19_conv_29[convPos_23_25_27_conv_naive_38] * conv_in[0+(0+wx_conv_naive_38)]
		tmp_25_27_conv_naive_38 = tmp_25_27_conv_naive_38 + computed_23_25_27_conv_naive_38
	}
	var computed_22_23_25_27_conv_naive_38 sint8 = kernel_19_conv_29[4] * conv_in[4]
	tmp_25_27_conv_naive_38 = tmp_25_27_conv_naive_38 + computed_22_23_25_27_conv_naive_38
	for wy_conv_naive_38 := 1; wy_conv_naive_38 < 4; wy_conv_naive_38 = wy_conv_naive_38 + 1 {
		var convPos_21_25_27_conv_naive_38 int = 0 + wy_conv_naive_38*5
		var computed_21_25_27_conv_naive_38 sint8 = kernel_19_conv_29[convPos_21_25_27_conv_naive_38] * conv_in[(0+wy_conv_naive_38)*28+(0)]
		tmp_25_27_conv_naive_38 = tmp_25_27_conv_naive_38 + computed_21_25_27_conv_naive_38
		for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
			var convPos_25_27_conv_naive_38 int = wx_conv_naive_38 + wy_conv_naive_38*5
			var computed_25_27_conv_naive_38 sint8 = kernel_19_conv_29[convPos_25_27_conv_naive_38] * conv_in[(0+wy_conv_naive_38)*28+(0+wx_conv_naive_38)]
			tmp_25_27_conv_naive_38 = tmp_25_27_conv_naive_38 + computed_25_27_conv_naive_38
		}
		var convPos_22_25_27_conv_naive_38 int = 4 + wy_conv_naive_38*5
		var computed_22_25_27_conv_naive_38 sint8 = kernel_19_conv_29[convPos_22_25_27_conv_naive_38] * conv_in[(0+wy_conv_naive_38)*28+(4)]
		tmp_25_27_conv_naive_38 = tmp_25_27_conv_naive_38 + computed_22_25_27_conv_naive_38
	}
	var computed_21_24_25_27_conv_naive_38 sint8 = kernel_19_conv_29[20] * conv_in[112]
	tmp_25_27_conv_naive_38 = tmp_25_27_conv_naive_38 + computed_21_24_25_27_conv_naive_38
	for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
		var convPos_24_25_27_conv_naive_38 int = wx_conv_naive_38 + 20
		var computed_24_25_27_conv_naive_38 sint8 = kernel_19_conv_29[convPos_24_25_27_conv_naive_38] * conv_in[112+(0+wx_conv_naive_38)]
		tmp_25_27_conv_naive_38 = tmp_25_27_conv_naive_38 + computed_24_25_27_conv_naive_38
	}
	var computed_22_24_25_27_conv_naive_38 sint8 = kernel_19_conv_29[24] * conv_in[116]
	tmp_25_27_conv_naive_38 = tmp_25_27_conv_naive_38 + computed_22_24_25_27_conv_naive_38
	out_conv_naive_38[0] = tmp_25_27_conv_naive_38
	for x_conv_naive_38 := 1; x_conv_naive_38 < 12; x_conv_naive_38 = x_conv_naive_38 + 1 {
		var oPos_27_conv_naive_38 int = x_conv_naive_38 + 0
		var tmp_27_conv_naive_38 sint8 = 0
		var computed_21_23_27_conv_naive_38 sint8 = kernel_19_conv_29[0] * conv_in[0+(x_conv_naive_38*2+0)]
		tmp_27_conv_naive_38 = tmp_27_conv_naive_38 + computed_21_23_27_conv_naive_38
		for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
			var convPos_23_27_conv_naive_38 int = wx_conv_naive_38 + 0
			var computed_23_27_conv_naive_38 sint8 = kernel_19_conv_29[convPos_23_27_conv_naive_38] * conv_in[0+(x_conv_naive_38*2+wx_conv_naive_38)]
			tmp_27_conv_naive_38 = tmp_27_conv_naive_38 + computed_23_27_conv_naive_38
		}
		var computed_22_23_27_conv_naive_38 sint8 = kernel_19_conv_29[4] * conv_in[0+(x_conv_naive_38*2+4)]
		tmp_27_conv_naive_38 = tmp_27_conv_naive_38 + computed_22_23_27_conv_naive_38
		for wy_conv_naive_38 := 1; wy_conv_naive_38 < 4; wy_conv_naive_38 = wy_conv_naive_38 + 1 {
			var convPos_21_27_conv_naive_38 int = 0 + wy_conv_naive_38*5
			var computed_21_27_conv_naive_38 sint8 = kernel_19_conv_29[convPos_21_27_conv_naive_38] * conv_in[(0+wy_conv_naive_38)*28+(x_conv_naive_38*2+0)]
			tmp_27_conv_naive_38 = tmp_27_conv_naive_38 + computed_21_27_conv_naive_38
			for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
				var convPos_27_conv_naive_38 int = wx_conv_naive_38 + wy_conv_naive_38*5
				var computed_27_conv_naive_38 sint8 = kernel_19_conv_29[convPos_27_conv_naive_38] * conv_in[(0+wy_conv_naive_38)*28+(x_conv_naive_38*2+wx_conv_naive_38)]
				tmp_27_conv_naive_38 = tmp_27_conv_naive_38 + computed_27_conv_naive_38
			}
			var convPos_22_27_conv_naive_38 int = 4 + wy_conv_naive_38*5
			var computed_22_27_conv_naive_38 sint8 = kernel_19_conv_29[convPos_22_27_conv_naive_38] * conv_in[(0+wy_conv_naive_38)*28+(x_conv_naive_38*2+4)]
			tmp_27_conv_naive_38 = tmp_27_conv_naive_38 + computed_22_27_conv_naive_38
		}
		var computed_21_24_27_conv_naive_38 sint8 = kernel_19_conv_29[20] * conv_in[112+(x_conv_naive_38*2+0)]
		tmp_27_conv_naive_38 = tmp_27_conv_naive_38 + computed_21_24_27_conv_naive_38
		for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
			var convPos_24_27_conv_naive_38 int = wx_conv_naive_38 + 20
			var computed_24_27_conv_naive_38 sint8 = kernel_19_conv_29[convPos_24_27_conv_naive_38] * conv_in[112+(x_conv_naive_38*2+wx_conv_naive_38)]
			tmp_27_conv_naive_38 = tmp_27_conv_naive_38 + computed_24_27_conv_naive_38
		}
		var computed_22_24_27_conv_naive_38 sint8 = kernel_19_conv_29[24] * conv_in[112+(x_conv_naive_38*2+4)]
		tmp_27_conv_naive_38 = tmp_27_conv_naive_38 + computed_22_24_27_conv_naive_38
		out_conv_naive_38[oPos_27_conv_naive_38] = tmp_27_conv_naive_38
	}
	var tmp_26_27_conv_naive_38 sint8 = 0
	var computed_21_23_26_27_conv_naive_38 sint8 = kernel_19_conv_29[0] * conv_in[24]
	tmp_26_27_conv_naive_38 = tmp_26_27_conv_naive_38 + computed_21_23_26_27_conv_naive_38
	for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
		var convPos_23_26_27_conv_naive_38 int = wx_conv_naive_38 + 0
		var computed_23_26_27_conv_naive_38 sint8 = kernel_19_conv_29[convPos_23_26_27_conv_naive_38] * conv_in[0+(24+wx_conv_naive_38)]
		tmp_26_27_conv_naive_38 = tmp_26_27_conv_naive_38 + computed_23_26_27_conv_naive_38
	}
	var computed_22_23_26_27_conv_naive_38 sint8 = kernel_19_conv_29[4] * conv_in[28]
	tmp_26_27_conv_naive_38 = tmp_26_27_conv_naive_38 + computed_22_23_26_27_conv_naive_38
	for wy_conv_naive_38 := 1; wy_conv_naive_38 < 4; wy_conv_naive_38 = wy_conv_naive_38 + 1 {
		var convPos_21_26_27_conv_naive_38 int = 0 + wy_conv_naive_38*5
		var computed_21_26_27_conv_naive_38 sint8 = kernel_19_conv_29[convPos_21_26_27_conv_naive_38] * conv_in[(0+wy_conv_naive_38)*28+(24)]
		tmp_26_27_conv_naive_38 = tmp_26_27_conv_naive_38 + computed_21_26_27_conv_naive_38
		for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
			var convPos_26_27_conv_naive_38 int = wx_conv_naive_38 + wy_conv_naive_38*5
			var computed_26_27_conv_naive_38 sint8 = kernel_19_conv_29[convPos_26_27_conv_naive_38] * conv_in[(0+wy_conv_naive_38)*28+(24+wx_conv_naive_38)]
			tmp_26_27_conv_naive_38 = tmp_26_27_conv_naive_38 + computed_26_27_conv_naive_38
		}
		var convPos_22_26_27_conv_naive_38 int = 4 + wy_conv_naive_38*5
		var computed_22_26_27_conv_naive_38 sint8 = kernel_19_conv_29[convPos_22_26_27_conv_naive_38] * conv_in[(0+wy_conv_naive_38)*28+(28)]
		tmp_26_27_conv_naive_38 = tmp_26_27_conv_naive_38 + computed_22_26_27_conv_naive_38
	}
	var computed_21_24_26_27_conv_naive_38 sint8 = kernel_19_conv_29[20] * conv_in[136]
	tmp_26_27_conv_naive_38 = tmp_26_27_conv_naive_38 + computed_21_24_26_27_conv_naive_38
	for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
		var convPos_24_26_27_conv_naive_38 int = wx_conv_naive_38 + 20
		var computed_24_26_27_conv_naive_38 sint8 = kernel_19_conv_29[convPos_24_26_27_conv_naive_38] * conv_in[112+(24+wx_conv_naive_38)]
		tmp_26_27_conv_naive_38 = tmp_26_27_conv_naive_38 + computed_24_26_27_conv_naive_38
	}
	var computed_22_24_26_27_conv_naive_38 sint8 = kernel_19_conv_29[24] * conv_in[140]
	tmp_26_27_conv_naive_38 = tmp_26_27_conv_naive_38 + computed_22_24_26_27_conv_naive_38
	out_conv_naive_38[12] = tmp_26_27_conv_naive_38
	for y_conv_naive_38 := 1; y_conv_naive_38 < 12; y_conv_naive_38 = y_conv_naive_38 + 1 {
		var oPos_25_conv_naive_38 int = 0 + y_conv_naive_38*13
		var tmp_25_conv_naive_38 sint8 = 0
		var computed_21_23_25_conv_naive_38 sint8 = kernel_19_conv_29[0] * conv_in[(y_conv_naive_38*2+0)*28+(0)]
		tmp_25_conv_naive_38 = tmp_25_conv_naive_38 + computed_21_23_25_conv_naive_38
		for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
			var convPos_23_25_conv_naive_38 int = wx_conv_naive_38 + 0
			var computed_23_25_conv_naive_38 sint8 = kernel_19_conv_29[convPos_23_25_conv_naive_38] * conv_in[(y_conv_naive_38*2+0)*28+(0+wx_conv_naive_38)]
			tmp_25_conv_naive_38 = tmp_25_conv_naive_38 + computed_23_25_conv_naive_38
		}
		var computed_22_23_25_conv_naive_38 sint8 = kernel_19_conv_29[4] * conv_in[(y_conv_naive_38*2+0)*28+(4)]
		tmp_25_conv_naive_38 = tmp_25_conv_naive_38 + computed_22_23_25_conv_naive_38
		for wy_conv_naive_38 := 1; wy_conv_naive_38 < 4; wy_conv_naive_38 = wy_conv_naive_38 + 1 {
			var convPos_21_25_conv_naive_38 int = 0 + wy_conv_naive_38*5
			var computed_21_25_conv_naive_38 sint8 = kernel_19_conv_29[convPos_21_25_conv_naive_38] * conv_in[(y_conv_naive_38*2+wy_conv_naive_38)*28+(0)]
			tmp_25_conv_naive_38 = tmp_25_conv_naive_38 + computed_21_25_conv_naive_38
			for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
				var convPos_25_conv_naive_38 int = wx_conv_naive_38 + wy_conv_naive_38*5
				var computed_25_conv_naive_38 sint8 = kernel_19_conv_29[convPos_25_conv_naive_38] * conv_in[(y_conv_naive_38*2+wy_conv_naive_38)*28+(0+wx_conv_naive_38)]
				tmp_25_conv_naive_38 = tmp_25_conv_naive_38 + computed_25_conv_naive_38
			}
			var convPos_22_25_conv_naive_38 int = 4 + wy_conv_naive_38*5
			var computed_22_25_conv_naive_38 sint8 = kernel_19_conv_29[convPos_22_25_conv_naive_38] * conv_in[(y_conv_naive_38*2+wy_conv_naive_38)*28+(4)]
			tmp_25_conv_naive_38 = tmp_25_conv_naive_38 + computed_22_25_conv_naive_38
		}
		var computed_21_24_25_conv_naive_38 sint8 = kernel_19_conv_29[20] * conv_in[(y_conv_naive_38*2+4)*28+(0)]
		tmp_25_conv_naive_38 = tmp_25_conv_naive_38 + computed_21_24_25_conv_naive_38
		for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
			var convPos_24_25_conv_naive_38 int = wx_conv_naive_38 + 20
			var computed_24_25_conv_naive_38 sint8 = kernel_19_conv_29[convPos_24_25_conv_naive_38] * conv_in[(y_conv_naive_38*2+4)*28+(0+wx_conv_naive_38)]
			tmp_25_conv_naive_38 = tmp_25_conv_naive_38 + computed_24_25_conv_naive_38
		}
		var computed_22_24_25_conv_naive_38 sint8 = kernel_19_conv_29[24] * conv_in[(y_conv_naive_38*2+4)*28+(4)]
		tmp_25_conv_naive_38 = tmp_25_conv_naive_38 + computed_22_24_25_conv_naive_38
		out_conv_naive_38[oPos_25_conv_naive_38] = tmp_25_conv_naive_38
		for x_conv_naive_38 := 1; x_conv_naive_38 < 12; x_conv_naive_38 = x_conv_naive_38 + 1 {
			var oPos_conv_naive_38 int = x_conv_naive_38 + y_conv_naive_38*13
			var tmp_conv_naive_38 sint8 = 0
			var computed_21_23_conv_naive_38 sint8 = kernel_19_conv_29[0] * conv_in[(y_conv_naive_38*2+0)*28+(x_conv_naive_38*2+0)]
			tmp_conv_naive_38 = tmp_conv_naive_38 + computed_21_23_conv_naive_38
			for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
				var convPos_23_conv_naive_38 int = wx_conv_naive_38 + 0
				var computed_23_conv_naive_38 sint8 = kernel_19_conv_29[convPos_23_conv_naive_38] * conv_in[(y_conv_naive_38*2+0)*28+(x_conv_naive_38*2+wx_conv_naive_38)]
				tmp_conv_naive_38 = tmp_conv_naive_38 + computed_23_conv_naive_38
			}
			var computed_22_23_conv_naive_38 sint8 = kernel_19_conv_29[4] * conv_in[(y_conv_naive_38*2+0)*28+(x_conv_naive_38*2+4)]
			tmp_conv_naive_38 = tmp_conv_naive_38 + computed_22_23_conv_naive_38
			for wy_conv_naive_38 := 1; wy_conv_naive_38 < 4; wy_conv_naive_38 = wy_conv_naive_38 + 1 {
				var convPos_21_conv_naive_38 int = 0 + wy_conv_naive_38*5
				var computed_21_conv_naive_38 sint8 = kernel_19_conv_29[convPos_21_conv_naive_38] * conv_in[(y_conv_naive_38*2+wy_conv_naive_38)*28+(x_conv_naive_38*2+0)]
				tmp_conv_naive_38 = tmp_conv_naive_38 + computed_21_conv_naive_38
				for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
					var convPos_conv_naive_38 int = wx_conv_naive_38 + wy_conv_naive_38*5
					var computed_conv_naive_38 sint8 = kernel_19_conv_29[convPos_conv_naive_38] * conv_in[(y_conv_naive_38*2+wy_conv_naive_38)*28+(x_conv_naive_38*2+wx_conv_naive_38)]
					tmp_conv_naive_38 = tmp_conv_naive_38 + computed_conv_naive_38
				}
				var convPos_22_conv_naive_38 int = 4 + wy_conv_naive_38*5
				var computed_22_conv_naive_38 sint8 = kernel_19_conv_29[convPos_22_conv_naive_38] * conv_in[(y_conv_naive_38*2+wy_conv_naive_38)*28+(x_conv_naive_38*2+4)]
				tmp_conv_naive_38 = tmp_conv_naive_38 + computed_22_conv_naive_38
			}
			var computed_21_24_conv_naive_38 sint8 = kernel_19_conv_29[20] * conv_in[(y_conv_naive_38*2+4)*28+(x_conv_naive_38*2+0)]
			tmp_conv_naive_38 = tmp_conv_naive_38 + computed_21_24_conv_naive_38
			for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
				var convPos_24_conv_naive_38 int = wx_conv_naive_38 + 20
				var computed_24_conv_naive_38 sint8 = kernel_19_conv_29[convPos_24_conv_naive_38] * conv_in[(y_conv_naive_38*2+4)*28+(x_conv_naive_38*2+wx_conv_naive_38)]
				tmp_conv_naive_38 = tmp_conv_naive_38 + computed_24_conv_naive_38
			}
			var computed_22_24_conv_naive_38 sint8 = kernel_19_conv_29[24] * conv_in[(y_conv_naive_38*2+4)*28+(x_conv_naive_38*2+4)]
			tmp_conv_naive_38 = tmp_conv_naive_38 + computed_22_24_conv_naive_38
			out_conv_naive_38[oPos_conv_naive_38] = tmp_conv_naive_38
		}
		var oPos_26_conv_naive_38 int = 12 + y_conv_naive_38*13
		var tmp_26_conv_naive_38 sint8 = 0
		var computed_21_23_26_conv_naive_38 sint8 = kernel_19_conv_29[0] * conv_in[(y_conv_naive_38*2+0)*28+(24)]
		tmp_26_conv_naive_38 = tmp_26_conv_naive_38 + computed_21_23_26_conv_naive_38
		for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
			var convPos_23_26_conv_naive_38 int = wx_conv_naive_38 + 0
			var computed_23_26_conv_naive_38 sint8 = kernel_19_conv_29[convPos_23_26_conv_naive_38] * conv_in[(y_conv_naive_38*2+0)*28+(24+wx_conv_naive_38)]
			tmp_26_conv_naive_38 = tmp_26_conv_naive_38 + computed_23_26_conv_naive_38
		}
		var computed_22_23_26_conv_naive_38 sint8 = kernel_19_conv_29[4] * conv_in[(y_conv_naive_38*2+0)*28+(28)]
		tmp_26_conv_naive_38 = tmp_26_conv_naive_38 + computed_22_23_26_conv_naive_38
		for wy_conv_naive_38 := 1; wy_conv_naive_38 < 4; wy_conv_naive_38 = wy_conv_naive_38 + 1 {
			var convPos_21_26_conv_naive_38 int = 0 + wy_conv_naive_38*5
			var computed_21_26_conv_naive_38 sint8 = kernel_19_conv_29[convPos_21_26_conv_naive_38] * conv_in[(y_conv_naive_38*2+wy_conv_naive_38)*28+(24)]
			tmp_26_conv_naive_38 = tmp_26_conv_naive_38 + computed_21_26_conv_naive_38
			for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
				var convPos_26_conv_naive_38 int = wx_conv_naive_38 + wy_conv_naive_38*5
				var computed_26_conv_naive_38 sint8 = kernel_19_conv_29[convPos_26_conv_naive_38] * conv_in[(y_conv_naive_38*2+wy_conv_naive_38)*28+(24+wx_conv_naive_38)]
				tmp_26_conv_naive_38 = tmp_26_conv_naive_38 + computed_26_conv_naive_38
			}
			var convPos_22_26_conv_naive_38 int = 4 + wy_conv_naive_38*5
			var computed_22_26_conv_naive_38 sint8 = kernel_19_conv_29[convPos_22_26_conv_naive_38] * conv_in[(y_conv_naive_38*2+wy_conv_naive_38)*28+(28)]
			tmp_26_conv_naive_38 = tmp_26_conv_naive_38 + computed_22_26_conv_naive_38
		}
		var computed_21_24_26_conv_naive_38 sint8 = kernel_19_conv_29[20] * conv_in[(y_conv_naive_38*2+4)*28+(24)]
		tmp_26_conv_naive_38 = tmp_26_conv_naive_38 + computed_21_24_26_conv_naive_38
		for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
			var convPos_24_26_conv_naive_38 int = wx_conv_naive_38 + 20
			var computed_24_26_conv_naive_38 sint8 = kernel_19_conv_29[convPos_24_26_conv_naive_38] * conv_in[(y_conv_naive_38*2+4)*28+(24+wx_conv_naive_38)]
			tmp_26_conv_naive_38 = tmp_26_conv_naive_38 + computed_24_26_conv_naive_38
		}
		var computed_22_24_26_conv_naive_38 sint8 = kernel_19_conv_29[24] * conv_in[(y_conv_naive_38*2+4)*28+(28)]
		tmp_26_conv_naive_38 = tmp_26_conv_naive_38 + computed_22_24_26_conv_naive_38
		out_conv_naive_38[oPos_26_conv_naive_38] = tmp_26_conv_naive_38
	}
	var tmp_25_28_conv_naive_38 sint8 = 0
	var computed_21_23_25_28_conv_naive_38 sint8 = kernel_19_conv_29[0] * conv_in[672]
	tmp_25_28_conv_naive_38 = tmp_25_28_conv_naive_38 + computed_21_23_25_28_conv_naive_38
	for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
		var convPos_23_25_28_conv_naive_38 int = wx_conv_naive_38 + 0
		var computed_23_25_28_conv_naive_38 sint8 = kernel_19_conv_29[convPos_23_25_28_conv_naive_38] * conv_in[672+(0+wx_conv_naive_38)]
		tmp_25_28_conv_naive_38 = tmp_25_28_conv_naive_38 + computed_23_25_28_conv_naive_38
	}
	var computed_22_23_25_28_conv_naive_38 sint8 = kernel_19_conv_29[4] * conv_in[676]
	tmp_25_28_conv_naive_38 = tmp_25_28_conv_naive_38 + computed_22_23_25_28_conv_naive_38
	for wy_conv_naive_38 := 1; wy_conv_naive_38 < 4; wy_conv_naive_38 = wy_conv_naive_38 + 1 {
		var convPos_21_25_28_conv_naive_38 int = 0 + wy_conv_naive_38*5
		var computed_21_25_28_conv_naive_38 sint8 = kernel_19_conv_29[convPos_21_25_28_conv_naive_38] * conv_in[(24+wy_conv_naive_38)*28+(0)]
		tmp_25_28_conv_naive_38 = tmp_25_28_conv_naive_38 + computed_21_25_28_conv_naive_38
		for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
			var convPos_25_28_conv_naive_38 int = wx_conv_naive_38 + wy_conv_naive_38*5
			var computed_25_28_conv_naive_38 sint8 = kernel_19_conv_29[convPos_25_28_conv_naive_38] * conv_in[(24+wy_conv_naive_38)*28+(0+wx_conv_naive_38)]
			tmp_25_28_conv_naive_38 = tmp_25_28_conv_naive_38 + computed_25_28_conv_naive_38
		}
		var convPos_22_25_28_conv_naive_38 int = 4 + wy_conv_naive_38*5
		var computed_22_25_28_conv_naive_38 sint8 = kernel_19_conv_29[convPos_22_25_28_conv_naive_38] * conv_in[(24+wy_conv_naive_38)*28+(4)]
		tmp_25_28_conv_naive_38 = tmp_25_28_conv_naive_38 + computed_22_25_28_conv_naive_38
	}
	var computed_21_24_25_28_conv_naive_38 sint8 = kernel_19_conv_29[20] * conv_in[784]
	tmp_25_28_conv_naive_38 = tmp_25_28_conv_naive_38 + computed_21_24_25_28_conv_naive_38
	for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
		var convPos_24_25_28_conv_naive_38 int = wx_conv_naive_38 + 20
		var computed_24_25_28_conv_naive_38 sint8 = kernel_19_conv_29[convPos_24_25_28_conv_naive_38] * conv_in[784+(0+wx_conv_naive_38)]
		tmp_25_28_conv_naive_38 = tmp_25_28_conv_naive_38 + computed_24_25_28_conv_naive_38
	}
	var computed_22_24_25_28_conv_naive_38 sint8 = kernel_19_conv_29[24] * conv_in[788]
	tmp_25_28_conv_naive_38 = tmp_25_28_conv_naive_38 + computed_22_24_25_28_conv_naive_38
	out_conv_naive_38[156] = tmp_25_28_conv_naive_38
	for x_conv_naive_38 := 1; x_conv_naive_38 < 12; x_conv_naive_38 = x_conv_naive_38 + 1 {
		var oPos_28_conv_naive_38 int = x_conv_naive_38 + 156
		var tmp_28_conv_naive_38 sint8 = 0
		var computed_21_23_28_conv_naive_38 sint8 = kernel_19_conv_29[0] * conv_in[672+(x_conv_naive_38*2+0)]
		tmp_28_conv_naive_38 = tmp_28_conv_naive_38 + computed_21_23_28_conv_naive_38
		for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
			var convPos_23_28_conv_naive_38 int = wx_conv_naive_38 + 0
			var computed_23_28_conv_naive_38 sint8 = kernel_19_conv_29[convPos_23_28_conv_naive_38] * conv_in[672+(x_conv_naive_38*2+wx_conv_naive_38)]
			tmp_28_conv_naive_38 = tmp_28_conv_naive_38 + computed_23_28_conv_naive_38
		}
		var computed_22_23_28_conv_naive_38 sint8 = kernel_19_conv_29[4] * conv_in[672+(x_conv_naive_38*2+4)]
		tmp_28_conv_naive_38 = tmp_28_conv_naive_38 + computed_22_23_28_conv_naive_38
		for wy_conv_naive_38 := 1; wy_conv_naive_38 < 4; wy_conv_naive_38 = wy_conv_naive_38 + 1 {
			var convPos_21_28_conv_naive_38 int = 0 + wy_conv_naive_38*5
			var computed_21_28_conv_naive_38 sint8 = kernel_19_conv_29[convPos_21_28_conv_naive_38] * conv_in[(24+wy_conv_naive_38)*28+(x_conv_naive_38*2+0)]
			tmp_28_conv_naive_38 = tmp_28_conv_naive_38 + computed_21_28_conv_naive_38
			for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
				var convPos_28_conv_naive_38 int = wx_conv_naive_38 + wy_conv_naive_38*5
				var computed_28_conv_naive_38 sint8 = kernel_19_conv_29[convPos_28_conv_naive_38] * conv_in[(24+wy_conv_naive_38)*28+(x_conv_naive_38*2+wx_conv_naive_38)]
				tmp_28_conv_naive_38 = tmp_28_conv_naive_38 + computed_28_conv_naive_38
			}
			var convPos_22_28_conv_naive_38 int = 4 + wy_conv_naive_38*5
			var computed_22_28_conv_naive_38 sint8 = kernel_19_conv_29[convPos_22_28_conv_naive_38] * conv_in[(24+wy_conv_naive_38)*28+(x_conv_naive_38*2+4)]
			tmp_28_conv_naive_38 = tmp_28_conv_naive_38 + computed_22_28_conv_naive_38
		}
		var computed_21_24_28_conv_naive_38 sint8 = kernel_19_conv_29[20] * conv_in[784+(x_conv_naive_38*2+0)]
		tmp_28_conv_naive_38 = tmp_28_conv_naive_38 + computed_21_24_28_conv_naive_38
		for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
			var convPos_24_28_conv_naive_38 int = wx_conv_naive_38 + 20
			var computed_24_28_conv_naive_38 sint8 = kernel_19_conv_29[convPos_24_28_conv_naive_38] * conv_in[784+(x_conv_naive_38*2+wx_conv_naive_38)]
			tmp_28_conv_naive_38 = tmp_28_conv_naive_38 + computed_24_28_conv_naive_38
		}
		var computed_22_24_28_conv_naive_38 sint8 = kernel_19_conv_29[24] * conv_in[784+(x_conv_naive_38*2+4)]
		tmp_28_conv_naive_38 = tmp_28_conv_naive_38 + computed_22_24_28_conv_naive_38
		out_conv_naive_38[oPos_28_conv_naive_38] = tmp_28_conv_naive_38
	}
	var tmp_26_28_conv_naive_38 sint8 = 0
	var computed_21_23_26_28_conv_naive_38 sint8 = kernel_19_conv_29[0] * conv_in[696]
	tmp_26_28_conv_naive_38 = tmp_26_28_conv_naive_38 + computed_21_23_26_28_conv_naive_38
	for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
		var convPos_23_26_28_conv_naive_38 int = wx_conv_naive_38 + 0
		var computed_23_26_28_conv_naive_38 sint8 = kernel_19_conv_29[convPos_23_26_28_conv_naive_38] * conv_in[672+(24+wx_conv_naive_38)]
		tmp_26_28_conv_naive_38 = tmp_26_28_conv_naive_38 + computed_23_26_28_conv_naive_38
	}
	var computed_22_23_26_28_conv_naive_38 sint8 = kernel_19_conv_29[4] * conv_in[700]
	tmp_26_28_conv_naive_38 = tmp_26_28_conv_naive_38 + computed_22_23_26_28_conv_naive_38
	for wy_conv_naive_38 := 1; wy_conv_naive_38 < 4; wy_conv_naive_38 = wy_conv_naive_38 + 1 {
		var convPos_21_26_28_conv_naive_38 int = 0 + wy_conv_naive_38*5
		var computed_21_26_28_conv_naive_38 sint8 = kernel_19_conv_29[convPos_21_26_28_conv_naive_38] * conv_in[(24+wy_conv_naive_38)*28+(24)]
		tmp_26_28_conv_naive_38 = tmp_26_28_conv_naive_38 + computed_21_26_28_conv_naive_38
		for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
			var convPos_26_28_conv_naive_38 int = wx_conv_naive_38 + wy_conv_naive_38*5
			var computed_26_28_conv_naive_38 sint8 = kernel_19_conv_29[convPos_26_28_conv_naive_38] * conv_in[(24+wy_conv_naive_38)*28+(24+wx_conv_naive_38)]
			tmp_26_28_conv_naive_38 = tmp_26_28_conv_naive_38 + computed_26_28_conv_naive_38
		}
		var convPos_22_26_28_conv_naive_38 int = 4 + wy_conv_naive_38*5
		var computed_22_26_28_conv_naive_38 sint8 = kernel_19_conv_29[convPos_22_26_28_conv_naive_38] * conv_in[(24+wy_conv_naive_38)*28+(28)]
		tmp_26_28_conv_naive_38 = tmp_26_28_conv_naive_38 + computed_22_26_28_conv_naive_38
	}
	var computed_21_24_26_28_conv_naive_38 sint8 = kernel_19_conv_29[20] * conv_in[808]
	tmp_26_28_conv_naive_38 = tmp_26_28_conv_naive_38 + computed_21_24_26_28_conv_naive_38
	for wx_conv_naive_38 := 1; wx_conv_naive_38 < 4; wx_conv_naive_38 = wx_conv_naive_38 + 1 {
		var convPos_24_26_28_conv_naive_38 int = wx_conv_naive_38 + 20
		var computed_24_26_28_conv_naive_38 sint8 = kernel_19_conv_29[convPos_24_26_28_conv_naive_38] * conv_in[784+(24+wx_conv_naive_38)]
		tmp_26_28_conv_naive_38 = tmp_26_28_conv_naive_38 + computed_24_26_28_conv_naive_38
	}
	var computed_22_24_26_28_conv_naive_38 sint8 = kernel_19_conv_29[24] * conv_in[812]
	tmp_26_28_conv_naive_38 = tmp_26_28_conv_naive_38 + computed_22_24_26_28_conv_naive_38
	out_conv_naive_38[168] = tmp_26_28_conv_naive_38
	var res_19_conv_29 []sint8 = out_conv_naive_38
	copy(output_conv_29[0:], res_19_conv_29)
	for i_conv_29 := 1; i_conv_29 < 4; i_conv_29 = i_conv_29 + 1 {
		var kernal_start_conv_29 int = i_conv_29 * 25
		var tmp_conv_29 int = i_conv_29 + 1
		var kernal_end_conv_29 int = tmp_conv_29 * 25
		var kernel_conv_29 []sint8 = kernel[kernal_start_conv_29:kernal_end_conv_29]
		var out_conv_naive_39 []sint8 = make([]sint8, 169)
		var tmp_25_27_conv_naive_39 sint8 = 0
		var computed_21_23_25_27_conv_naive_39 sint8 = kernel_conv_29[0] * conv_in[0]
		tmp_25_27_conv_naive_39 = tmp_25_27_conv_naive_39 + computed_21_23_25_27_conv_naive_39
		for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
			var convPos_23_25_27_conv_naive_39 int = wx_conv_naive_39 + 0
			var computed_23_25_27_conv_naive_39 sint8 = kernel_conv_29[convPos_23_25_27_conv_naive_39] * conv_in[0+(0+wx_conv_naive_39)]
			tmp_25_27_conv_naive_39 = tmp_25_27_conv_naive_39 + computed_23_25_27_conv_naive_39
		}
		var computed_22_23_25_27_conv_naive_39 sint8 = kernel_conv_29[4] * conv_in[4]
		tmp_25_27_conv_naive_39 = tmp_25_27_conv_naive_39 + computed_22_23_25_27_conv_naive_39
		for wy_conv_naive_39 := 1; wy_conv_naive_39 < 4; wy_conv_naive_39 = wy_conv_naive_39 + 1 {
			var convPos_21_25_27_conv_naive_39 int = 0 + wy_conv_naive_39*5
			var computed_21_25_27_conv_naive_39 sint8 = kernel_conv_29[convPos_21_25_27_conv_naive_39] * conv_in[(0+wy_conv_naive_39)*28+(0)]
			tmp_25_27_conv_naive_39 = tmp_25_27_conv_naive_39 + computed_21_25_27_conv_naive_39
			for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
				var convPos_25_27_conv_naive_39 int = wx_conv_naive_39 + wy_conv_naive_39*5
				var computed_25_27_conv_naive_39 sint8 = kernel_conv_29[convPos_25_27_conv_naive_39] * conv_in[(0+wy_conv_naive_39)*28+(0+wx_conv_naive_39)]
				tmp_25_27_conv_naive_39 = tmp_25_27_conv_naive_39 + computed_25_27_conv_naive_39
			}
			var convPos_22_25_27_conv_naive_39 int = 4 + wy_conv_naive_39*5
			var computed_22_25_27_conv_naive_39 sint8 = kernel_conv_29[convPos_22_25_27_conv_naive_39] * conv_in[(0+wy_conv_naive_39)*28+(4)]
			tmp_25_27_conv_naive_39 = tmp_25_27_conv_naive_39 + computed_22_25_27_conv_naive_39
		}
		var computed_21_24_25_27_conv_naive_39 sint8 = kernel_conv_29[20] * conv_in[112]
		tmp_25_27_conv_naive_39 = tmp_25_27_conv_naive_39 + computed_21_24_25_27_conv_naive_39
		for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
			var convPos_24_25_27_conv_naive_39 int = wx_conv_naive_39 + 20
			var computed_24_25_27_conv_naive_39 sint8 = kernel_conv_29[convPos_24_25_27_conv_naive_39] * conv_in[112+(0+wx_conv_naive_39)]
			tmp_25_27_conv_naive_39 = tmp_25_27_conv_naive_39 + computed_24_25_27_conv_naive_39
		}
		var computed_22_24_25_27_conv_naive_39 sint8 = kernel_conv_29[24] * conv_in[116]
		tmp_25_27_conv_naive_39 = tmp_25_27_conv_naive_39 + computed_22_24_25_27_conv_naive_39
		out_conv_naive_39[0] = tmp_25_27_conv_naive_39
		for x_conv_naive_39 := 1; x_conv_naive_39 < 12; x_conv_naive_39 = x_conv_naive_39 + 1 {
			var oPos_27_conv_naive_39 int = x_conv_naive_39 + 0
			var tmp_27_conv_naive_39 sint8 = 0
			var computed_21_23_27_conv_naive_39 sint8 = kernel_conv_29[0] * conv_in[0+(x_conv_naive_39*2+0)]
			tmp_27_conv_naive_39 = tmp_27_conv_naive_39 + computed_21_23_27_conv_naive_39
			for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
				var convPos_23_27_conv_naive_39 int = wx_conv_naive_39 + 0
				var computed_23_27_conv_naive_39 sint8 = kernel_conv_29[convPos_23_27_conv_naive_39] * conv_in[0+(x_conv_naive_39*2+wx_conv_naive_39)]
				tmp_27_conv_naive_39 = tmp_27_conv_naive_39 + computed_23_27_conv_naive_39
			}
			var computed_22_23_27_conv_naive_39 sint8 = kernel_conv_29[4] * conv_in[0+(x_conv_naive_39*2+4)]
			tmp_27_conv_naive_39 = tmp_27_conv_naive_39 + computed_22_23_27_conv_naive_39
			for wy_conv_naive_39 := 1; wy_conv_naive_39 < 4; wy_conv_naive_39 = wy_conv_naive_39 + 1 {
				var convPos_21_27_conv_naive_39 int = 0 + wy_conv_naive_39*5
				var computed_21_27_conv_naive_39 sint8 = kernel_conv_29[convPos_21_27_conv_naive_39] * conv_in[(0+wy_conv_naive_39)*28+(x_conv_naive_39*2+0)]
				tmp_27_conv_naive_39 = tmp_27_conv_naive_39 + computed_21_27_conv_naive_39
				for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
					var convPos_27_conv_naive_39 int = wx_conv_naive_39 + wy_conv_naive_39*5
					var computed_27_conv_naive_39 sint8 = kernel_conv_29[convPos_27_conv_naive_39] * conv_in[(0+wy_conv_naive_39)*28+(x_conv_naive_39*2+wx_conv_naive_39)]
					tmp_27_conv_naive_39 = tmp_27_conv_naive_39 + computed_27_conv_naive_39
				}
				var convPos_22_27_conv_naive_39 int = 4 + wy_conv_naive_39*5
				var computed_22_27_conv_naive_39 sint8 = kernel_conv_29[convPos_22_27_conv_naive_39] * conv_in[(0+wy_conv_naive_39)*28+(x_conv_naive_39*2+4)]
				tmp_27_conv_naive_39 = tmp_27_conv_naive_39 + computed_22_27_conv_naive_39
			}
			var computed_21_24_27_conv_naive_39 sint8 = kernel_conv_29[20] * conv_in[112+(x_conv_naive_39*2+0)]
			tmp_27_conv_naive_39 = tmp_27_conv_naive_39 + computed_21_24_27_conv_naive_39
			for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
				var convPos_24_27_conv_naive_39 int = wx_conv_naive_39 + 20
				var computed_24_27_conv_naive_39 sint8 = kernel_conv_29[convPos_24_27_conv_naive_39] * conv_in[112+(x_conv_naive_39*2+wx_conv_naive_39)]
				tmp_27_conv_naive_39 = tmp_27_conv_naive_39 + computed_24_27_conv_naive_39
			}
			var computed_22_24_27_conv_naive_39 sint8 = kernel_conv_29[24] * conv_in[112+(x_conv_naive_39*2+4)]
			tmp_27_conv_naive_39 = tmp_27_conv_naive_39 + computed_22_24_27_conv_naive_39
			out_conv_naive_39[oPos_27_conv_naive_39] = tmp_27_conv_naive_39
		}
		var tmp_26_27_conv_naive_39 sint8 = 0
		var computed_21_23_26_27_conv_naive_39 sint8 = kernel_conv_29[0] * conv_in[24]
		tmp_26_27_conv_naive_39 = tmp_26_27_conv_naive_39 + computed_21_23_26_27_conv_naive_39
		for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
			var convPos_23_26_27_conv_naive_39 int = wx_conv_naive_39 + 0
			var computed_23_26_27_conv_naive_39 sint8 = kernel_conv_29[convPos_23_26_27_conv_naive_39] * conv_in[0+(24+wx_conv_naive_39)]
			tmp_26_27_conv_naive_39 = tmp_26_27_conv_naive_39 + computed_23_26_27_conv_naive_39
		}
		var computed_22_23_26_27_conv_naive_39 sint8 = kernel_conv_29[4] * conv_in[28]
		tmp_26_27_conv_naive_39 = tmp_26_27_conv_naive_39 + computed_22_23_26_27_conv_naive_39
		for wy_conv_naive_39 := 1; wy_conv_naive_39 < 4; wy_conv_naive_39 = wy_conv_naive_39 + 1 {
			var convPos_21_26_27_conv_naive_39 int = 0 + wy_conv_naive_39*5
			var computed_21_26_27_conv_naive_39 sint8 = kernel_conv_29[convPos_21_26_27_conv_naive_39] * conv_in[(0+wy_conv_naive_39)*28+(24)]
			tmp_26_27_conv_naive_39 = tmp_26_27_conv_naive_39 + computed_21_26_27_conv_naive_39
			for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
				var convPos_26_27_conv_naive_39 int = wx_conv_naive_39 + wy_conv_naive_39*5
				var computed_26_27_conv_naive_39 sint8 = kernel_conv_29[convPos_26_27_conv_naive_39] * conv_in[(0+wy_conv_naive_39)*28+(24+wx_conv_naive_39)]
				tmp_26_27_conv_naive_39 = tmp_26_27_conv_naive_39 + computed_26_27_conv_naive_39
			}
			var convPos_22_26_27_conv_naive_39 int = 4 + wy_conv_naive_39*5
			var computed_22_26_27_conv_naive_39 sint8 = kernel_conv_29[convPos_22_26_27_conv_naive_39] * conv_in[(0+wy_conv_naive_39)*28+(28)]
			tmp_26_27_conv_naive_39 = tmp_26_27_conv_naive_39 + computed_22_26_27_conv_naive_39
		}
		var computed_21_24_26_27_conv_naive_39 sint8 = kernel_conv_29[20] * conv_in[136]
		tmp_26_27_conv_naive_39 = tmp_26_27_conv_naive_39 + computed_21_24_26_27_conv_naive_39
		for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
			var convPos_24_26_27_conv_naive_39 int = wx_conv_naive_39 + 20
			var computed_24_26_27_conv_naive_39 sint8 = kernel_conv_29[convPos_24_26_27_conv_naive_39] * conv_in[112+(24+wx_conv_naive_39)]
			tmp_26_27_conv_naive_39 = tmp_26_27_conv_naive_39 + computed_24_26_27_conv_naive_39
		}
		var computed_22_24_26_27_conv_naive_39 sint8 = kernel_conv_29[24] * conv_in[140]
		tmp_26_27_conv_naive_39 = tmp_26_27_conv_naive_39 + computed_22_24_26_27_conv_naive_39
		out_conv_naive_39[12] = tmp_26_27_conv_naive_39
		for y_conv_naive_39 := 1; y_conv_naive_39 < 12; y_conv_naive_39 = y_conv_naive_39 + 1 {
			var oPos_25_conv_naive_39 int = 0 + y_conv_naive_39*13
			var tmp_25_conv_naive_39 sint8 = 0
			var computed_21_23_25_conv_naive_39 sint8 = kernel_conv_29[0] * conv_in[(y_conv_naive_39*2+0)*28+(0)]
			tmp_25_conv_naive_39 = tmp_25_conv_naive_39 + computed_21_23_25_conv_naive_39
			for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
				var convPos_23_25_conv_naive_39 int = wx_conv_naive_39 + 0
				var computed_23_25_conv_naive_39 sint8 = kernel_conv_29[convPos_23_25_conv_naive_39] * conv_in[(y_conv_naive_39*2+0)*28+(0+wx_conv_naive_39)]
				tmp_25_conv_naive_39 = tmp_25_conv_naive_39 + computed_23_25_conv_naive_39
			}
			var computed_22_23_25_conv_naive_39 sint8 = kernel_conv_29[4] * conv_in[(y_conv_naive_39*2+0)*28+(4)]
			tmp_25_conv_naive_39 = tmp_25_conv_naive_39 + computed_22_23_25_conv_naive_39
			for wy_conv_naive_39 := 1; wy_conv_naive_39 < 4; wy_conv_naive_39 = wy_conv_naive_39 + 1 {
				var convPos_21_25_conv_naive_39 int = 0 + wy_conv_naive_39*5
				var computed_21_25_conv_naive_39 sint8 = kernel_conv_29[convPos_21_25_conv_naive_39] * conv_in[(y_conv_naive_39*2+wy_conv_naive_39)*28+(0)]
				tmp_25_conv_naive_39 = tmp_25_conv_naive_39 + computed_21_25_conv_naive_39
				for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
					var convPos_25_conv_naive_39 int = wx_conv_naive_39 + wy_conv_naive_39*5
					var computed_25_conv_naive_39 sint8 = kernel_conv_29[convPos_25_conv_naive_39] * conv_in[(y_conv_naive_39*2+wy_conv_naive_39)*28+(0+wx_conv_naive_39)]
					tmp_25_conv_naive_39 = tmp_25_conv_naive_39 + computed_25_conv_naive_39
				}
				var convPos_22_25_conv_naive_39 int = 4 + wy_conv_naive_39*5
				var computed_22_25_conv_naive_39 sint8 = kernel_conv_29[convPos_22_25_conv_naive_39] * conv_in[(y_conv_naive_39*2+wy_conv_naive_39)*28+(4)]
				tmp_25_conv_naive_39 = tmp_25_conv_naive_39 + computed_22_25_conv_naive_39
			}
			var computed_21_24_25_conv_naive_39 sint8 = kernel_conv_29[20] * conv_in[(y_conv_naive_39*2+4)*28+(0)]
			tmp_25_conv_naive_39 = tmp_25_conv_naive_39 + computed_21_24_25_conv_naive_39
			for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
				var convPos_24_25_conv_naive_39 int = wx_conv_naive_39 + 20
				var computed_24_25_conv_naive_39 sint8 = kernel_conv_29[convPos_24_25_conv_naive_39] * conv_in[(y_conv_naive_39*2+4)*28+(0+wx_conv_naive_39)]
				tmp_25_conv_naive_39 = tmp_25_conv_naive_39 + computed_24_25_conv_naive_39
			}
			var computed_22_24_25_conv_naive_39 sint8 = kernel_conv_29[24] * conv_in[(y_conv_naive_39*2+4)*28+(4)]
			tmp_25_conv_naive_39 = tmp_25_conv_naive_39 + computed_22_24_25_conv_naive_39
			out_conv_naive_39[oPos_25_conv_naive_39] = tmp_25_conv_naive_39
			for x_conv_naive_39 := 1; x_conv_naive_39 < 12; x_conv_naive_39 = x_conv_naive_39 + 1 {
				var oPos_conv_naive_39 int = x_conv_naive_39 + y_conv_naive_39*13
				var tmp_conv_naive_39 sint8 = 0
				var computed_21_23_conv_naive_39 sint8 = kernel_conv_29[0] * conv_in[(y_conv_naive_39*2+0)*28+(x_conv_naive_39*2+0)]
				tmp_conv_naive_39 = tmp_conv_naive_39 + computed_21_23_conv_naive_39
				for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
					var convPos_23_conv_naive_39 int = wx_conv_naive_39 + 0
					var computed_23_conv_naive_39 sint8 = kernel_conv_29[convPos_23_conv_naive_39] * conv_in[(y_conv_naive_39*2+0)*28+(x_conv_naive_39*2+wx_conv_naive_39)]
					tmp_conv_naive_39 = tmp_conv_naive_39 + computed_23_conv_naive_39
				}
				var computed_22_23_conv_naive_39 sint8 = kernel_conv_29[4] * conv_in[(y_conv_naive_39*2+0)*28+(x_conv_naive_39*2+4)]
				tmp_conv_naive_39 = tmp_conv_naive_39 + computed_22_23_conv_naive_39
				for wy_conv_naive_39 := 1; wy_conv_naive_39 < 4; wy_conv_naive_39 = wy_conv_naive_39 + 1 {
					var convPos_21_conv_naive_39 int = 0 + wy_conv_naive_39*5
					var computed_21_conv_naive_39 sint8 = kernel_conv_29[convPos_21_conv_naive_39] * conv_in[(y_conv_naive_39*2+wy_conv_naive_39)*28+(x_conv_naive_39*2+0)]
					tmp_conv_naive_39 = tmp_conv_naive_39 + computed_21_conv_naive_39
					for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
						var convPos_conv_naive_39 int = wx_conv_naive_39 + wy_conv_naive_39*5
						var computed_conv_naive_39 sint8 = kernel_conv_29[convPos_conv_naive_39] * conv_in[(y_conv_naive_39*2+wy_conv_naive_39)*28+(x_conv_naive_39*2+wx_conv_naive_39)]
						tmp_conv_naive_39 = tmp_conv_naive_39 + computed_conv_naive_39
					}
					var convPos_22_conv_naive_39 int = 4 + wy_conv_naive_39*5
					var computed_22_conv_naive_39 sint8 = kernel_conv_29[convPos_22_conv_naive_39] * conv_in[(y_conv_naive_39*2+wy_conv_naive_39)*28+(x_conv_naive_39*2+4)]
					tmp_conv_naive_39 = tmp_conv_naive_39 + computed_22_conv_naive_39
				}
				var computed_21_24_conv_naive_39 sint8 = kernel_conv_29[20] * conv_in[(y_conv_naive_39*2+4)*28+(x_conv_naive_39*2+0)]
				tmp_conv_naive_39 = tmp_conv_naive_39 + computed_21_24_conv_naive_39
				for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
					var convPos_24_conv_naive_39 int = wx_conv_naive_39 + 20
					var computed_24_conv_naive_39 sint8 = kernel_conv_29[convPos_24_conv_naive_39] * conv_in[(y_conv_naive_39*2+4)*28+(x_conv_naive_39*2+wx_conv_naive_39)]
					tmp_conv_naive_39 = tmp_conv_naive_39 + computed_24_conv_naive_39
				}
				var computed_22_24_conv_naive_39 sint8 = kernel_conv_29[24] * conv_in[(y_conv_naive_39*2+4)*28+(x_conv_naive_39*2+4)]
				tmp_conv_naive_39 = tmp_conv_naive_39 + computed_22_24_conv_naive_39
				out_conv_naive_39[oPos_conv_naive_39] = tmp_conv_naive_39
			}
			var oPos_26_conv_naive_39 int = 12 + y_conv_naive_39*13
			var tmp_26_conv_naive_39 sint8 = 0
			var computed_21_23_26_conv_naive_39 sint8 = kernel_conv_29[0] * conv_in[(y_conv_naive_39*2+0)*28+(24)]
			tmp_26_conv_naive_39 = tmp_26_conv_naive_39 + computed_21_23_26_conv_naive_39
			for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
				var convPos_23_26_conv_naive_39 int = wx_conv_naive_39 + 0
				var computed_23_26_conv_naive_39 sint8 = kernel_conv_29[convPos_23_26_conv_naive_39] * conv_in[(y_conv_naive_39*2+0)*28+(24+wx_conv_naive_39)]
				tmp_26_conv_naive_39 = tmp_26_conv_naive_39 + computed_23_26_conv_naive_39
			}
			var computed_22_23_26_conv_naive_39 sint8 = kernel_conv_29[4] * conv_in[(y_conv_naive_39*2+0)*28+(28)]
			tmp_26_conv_naive_39 = tmp_26_conv_naive_39 + computed_22_23_26_conv_naive_39
			for wy_conv_naive_39 := 1; wy_conv_naive_39 < 4; wy_conv_naive_39 = wy_conv_naive_39 + 1 {
				var convPos_21_26_conv_naive_39 int = 0 + wy_conv_naive_39*5
				var computed_21_26_conv_naive_39 sint8 = kernel_conv_29[convPos_21_26_conv_naive_39] * conv_in[(y_conv_naive_39*2+wy_conv_naive_39)*28+(24)]
				tmp_26_conv_naive_39 = tmp_26_conv_naive_39 + computed_21_26_conv_naive_39
				for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
					var convPos_26_conv_naive_39 int = wx_conv_naive_39 + wy_conv_naive_39*5
					var computed_26_conv_naive_39 sint8 = kernel_conv_29[convPos_26_conv_naive_39] * conv_in[(y_conv_naive_39*2+wy_conv_naive_39)*28+(24+wx_conv_naive_39)]
					tmp_26_conv_naive_39 = tmp_26_conv_naive_39 + computed_26_conv_naive_39
				}
				var convPos_22_26_conv_naive_39 int = 4 + wy_conv_naive_39*5
				var computed_22_26_conv_naive_39 sint8 = kernel_conv_29[convPos_22_26_conv_naive_39] * conv_in[(y_conv_naive_39*2+wy_conv_naive_39)*28+(28)]
				tmp_26_conv_naive_39 = tmp_26_conv_naive_39 + computed_22_26_conv_naive_39
			}
			var computed_21_24_26_conv_naive_39 sint8 = kernel_conv_29[20] * conv_in[(y_conv_naive_39*2+4)*28+(24)]
			tmp_26_conv_naive_39 = tmp_26_conv_naive_39 + computed_21_24_26_conv_naive_39
			for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
				var convPos_24_26_conv_naive_39 int = wx_conv_naive_39 + 20
				var computed_24_26_conv_naive_39 sint8 = kernel_conv_29[convPos_24_26_conv_naive_39] * conv_in[(y_conv_naive_39*2+4)*28+(24+wx_conv_naive_39)]
				tmp_26_conv_naive_39 = tmp_26_conv_naive_39 + computed_24_26_conv_naive_39
			}
			var computed_22_24_26_conv_naive_39 sint8 = kernel_conv_29[24] * conv_in[(y_conv_naive_39*2+4)*28+(28)]
			tmp_26_conv_naive_39 = tmp_26_conv_naive_39 + computed_22_24_26_conv_naive_39
			out_conv_naive_39[oPos_26_conv_naive_39] = tmp_26_conv_naive_39
		}
		var tmp_25_28_conv_naive_39 sint8 = 0
		var computed_21_23_25_28_conv_naive_39 sint8 = kernel_conv_29[0] * conv_in[672]
		tmp_25_28_conv_naive_39 = tmp_25_28_conv_naive_39 + computed_21_23_25_28_conv_naive_39
		for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
			var convPos_23_25_28_conv_naive_39 int = wx_conv_naive_39 + 0
			var computed_23_25_28_conv_naive_39 sint8 = kernel_conv_29[convPos_23_25_28_conv_naive_39] * conv_in[672+(0+wx_conv_naive_39)]
			tmp_25_28_conv_naive_39 = tmp_25_28_conv_naive_39 + computed_23_25_28_conv_naive_39
		}
		var computed_22_23_25_28_conv_naive_39 sint8 = kernel_conv_29[4] * conv_in[676]
		tmp_25_28_conv_naive_39 = tmp_25_28_conv_naive_39 + computed_22_23_25_28_conv_naive_39
		for wy_conv_naive_39 := 1; wy_conv_naive_39 < 4; wy_conv_naive_39 = wy_conv_naive_39 + 1 {
			var convPos_21_25_28_conv_naive_39 int = 0 + wy_conv_naive_39*5
			var computed_21_25_28_conv_naive_39 sint8 = kernel_conv_29[convPos_21_25_28_conv_naive_39] * conv_in[(24+wy_conv_naive_39)*28+(0)]
			tmp_25_28_conv_naive_39 = tmp_25_28_conv_naive_39 + computed_21_25_28_conv_naive_39
			for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
				var convPos_25_28_conv_naive_39 int = wx_conv_naive_39 + wy_conv_naive_39*5
				var computed_25_28_conv_naive_39 sint8 = kernel_conv_29[convPos_25_28_conv_naive_39] * conv_in[(24+wy_conv_naive_39)*28+(0+wx_conv_naive_39)]
				tmp_25_28_conv_naive_39 = tmp_25_28_conv_naive_39 + computed_25_28_conv_naive_39
			}
			var convPos_22_25_28_conv_naive_39 int = 4 + wy_conv_naive_39*5
			var computed_22_25_28_conv_naive_39 sint8 = kernel_conv_29[convPos_22_25_28_conv_naive_39] * conv_in[(24+wy_conv_naive_39)*28+(4)]
			tmp_25_28_conv_naive_39 = tmp_25_28_conv_naive_39 + computed_22_25_28_conv_naive_39
		}
		var computed_21_24_25_28_conv_naive_39 sint8 = kernel_conv_29[20] * conv_in[784]
		tmp_25_28_conv_naive_39 = tmp_25_28_conv_naive_39 + computed_21_24_25_28_conv_naive_39
		for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
			var convPos_24_25_28_conv_naive_39 int = wx_conv_naive_39 + 20
			var computed_24_25_28_conv_naive_39 sint8 = kernel_conv_29[convPos_24_25_28_conv_naive_39] * conv_in[784+(0+wx_conv_naive_39)]
			tmp_25_28_conv_naive_39 = tmp_25_28_conv_naive_39 + computed_24_25_28_conv_naive_39
		}
		var computed_22_24_25_28_conv_naive_39 sint8 = kernel_conv_29[24] * conv_in[788]
		tmp_25_28_conv_naive_39 = tmp_25_28_conv_naive_39 + computed_22_24_25_28_conv_naive_39
		out_conv_naive_39[156] = tmp_25_28_conv_naive_39
		for x_conv_naive_39 := 1; x_conv_naive_39 < 12; x_conv_naive_39 = x_conv_naive_39 + 1 {
			var oPos_28_conv_naive_39 int = x_conv_naive_39 + 156
			var tmp_28_conv_naive_39 sint8 = 0
			var computed_21_23_28_conv_naive_39 sint8 = kernel_conv_29[0] * conv_in[672+(x_conv_naive_39*2+0)]
			tmp_28_conv_naive_39 = tmp_28_conv_naive_39 + computed_21_23_28_conv_naive_39
			for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
				var convPos_23_28_conv_naive_39 int = wx_conv_naive_39 + 0
				var computed_23_28_conv_naive_39 sint8 = kernel_conv_29[convPos_23_28_conv_naive_39] * conv_in[672+(x_conv_naive_39*2+wx_conv_naive_39)]
				tmp_28_conv_naive_39 = tmp_28_conv_naive_39 + computed_23_28_conv_naive_39
			}
			var computed_22_23_28_conv_naive_39 sint8 = kernel_conv_29[4] * conv_in[672+(x_conv_naive_39*2+4)]
			tmp_28_conv_naive_39 = tmp_28_conv_naive_39 + computed_22_23_28_conv_naive_39
			for wy_conv_naive_39 := 1; wy_conv_naive_39 < 4; wy_conv_naive_39 = wy_conv_naive_39 + 1 {
				var convPos_21_28_conv_naive_39 int = 0 + wy_conv_naive_39*5
				var computed_21_28_conv_naive_39 sint8 = kernel_conv_29[convPos_21_28_conv_naive_39] * conv_in[(24+wy_conv_naive_39)*28+(x_conv_naive_39*2+0)]
				tmp_28_conv_naive_39 = tmp_28_conv_naive_39 + computed_21_28_conv_naive_39
				for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
					var convPos_28_conv_naive_39 int = wx_conv_naive_39 + wy_conv_naive_39*5
					var computed_28_conv_naive_39 sint8 = kernel_conv_29[convPos_28_conv_naive_39] * conv_in[(24+wy_conv_naive_39)*28+(x_conv_naive_39*2+wx_conv_naive_39)]
					tmp_28_conv_naive_39 = tmp_28_conv_naive_39 + computed_28_conv_naive_39
				}
				var convPos_22_28_conv_naive_39 int = 4 + wy_conv_naive_39*5
				var computed_22_28_conv_naive_39 sint8 = kernel_conv_29[convPos_22_28_conv_naive_39] * conv_in[(24+wy_conv_naive_39)*28+(x_conv_naive_39*2+4)]
				tmp_28_conv_naive_39 = tmp_28_conv_naive_39 + computed_22_28_conv_naive_39
			}
			var computed_21_24_28_conv_naive_39 sint8 = kernel_conv_29[20] * conv_in[784+(x_conv_naive_39*2+0)]
			tmp_28_conv_naive_39 = tmp_28_conv_naive_39 + computed_21_24_28_conv_naive_39
			for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
				var convPos_24_28_conv_naive_39 int = wx_conv_naive_39 + 20
				var computed_24_28_conv_naive_39 sint8 = kernel_conv_29[convPos_24_28_conv_naive_39] * conv_in[784+(x_conv_naive_39*2+wx_conv_naive_39)]
				tmp_28_conv_naive_39 = tmp_28_conv_naive_39 + computed_24_28_conv_naive_39
			}
			var computed_22_24_28_conv_naive_39 sint8 = kernel_conv_29[24] * conv_in[784+(x_conv_naive_39*2+4)]
			tmp_28_conv_naive_39 = tmp_28_conv_naive_39 + computed_22_24_28_conv_naive_39
			out_conv_naive_39[oPos_28_conv_naive_39] = tmp_28_conv_naive_39
		}
		var tmp_26_28_conv_naive_39 sint8 = 0
		var computed_21_23_26_28_conv_naive_39 sint8 = kernel_conv_29[0] * conv_in[696]
		tmp_26_28_conv_naive_39 = tmp_26_28_conv_naive_39 + computed_21_23_26_28_conv_naive_39
		for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
			var convPos_23_26_28_conv_naive_39 int = wx_conv_naive_39 + 0
			var computed_23_26_28_conv_naive_39 sint8 = kernel_conv_29[convPos_23_26_28_conv_naive_39] * conv_in[672+(24+wx_conv_naive_39)]
			tmp_26_28_conv_naive_39 = tmp_26_28_conv_naive_39 + computed_23_26_28_conv_naive_39
		}
		var computed_22_23_26_28_conv_naive_39 sint8 = kernel_conv_29[4] * conv_in[700]
		tmp_26_28_conv_naive_39 = tmp_26_28_conv_naive_39 + computed_22_23_26_28_conv_naive_39
		for wy_conv_naive_39 := 1; wy_conv_naive_39 < 4; wy_conv_naive_39 = wy_conv_naive_39 + 1 {
			var convPos_21_26_28_conv_naive_39 int = 0 + wy_conv_naive_39*5
			var computed_21_26_28_conv_naive_39 sint8 = kernel_conv_29[convPos_21_26_28_conv_naive_39] * conv_in[(24+wy_conv_naive_39)*28+(24)]
			tmp_26_28_conv_naive_39 = tmp_26_28_conv_naive_39 + computed_21_26_28_conv_naive_39
			for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
				var convPos_26_28_conv_naive_39 int = wx_conv_naive_39 + wy_conv_naive_39*5
				var computed_26_28_conv_naive_39 sint8 = kernel_conv_29[convPos_26_28_conv_naive_39] * conv_in[(24+wy_conv_naive_39)*28+(24+wx_conv_naive_39)]
				tmp_26_28_conv_naive_39 = tmp_26_28_conv_naive_39 + computed_26_28_conv_naive_39
			}
			var convPos_22_26_28_conv_naive_39 int = 4 + wy_conv_naive_39*5
			var computed_22_26_28_conv_naive_39 sint8 = kernel_conv_29[convPos_22_26_28_conv_naive_39] * conv_in[(24+wy_conv_naive_39)*28+(28)]
			tmp_26_28_conv_naive_39 = tmp_26_28_conv_naive_39 + computed_22_26_28_conv_naive_39
		}
		var computed_21_24_26_28_conv_naive_39 sint8 = kernel_conv_29[20] * conv_in[808]
		tmp_26_28_conv_naive_39 = tmp_26_28_conv_naive_39 + computed_21_24_26_28_conv_naive_39
		for wx_conv_naive_39 := 1; wx_conv_naive_39 < 4; wx_conv_naive_39 = wx_conv_naive_39 + 1 {
			var convPos_24_26_28_conv_naive_39 int = wx_conv_naive_39 + 20
			var computed_24_26_28_conv_naive_39 sint8 = kernel_conv_29[convPos_24_26_28_conv_naive_39] * conv_in[784+(24+wx_conv_naive_39)]
			tmp_26_28_conv_naive_39 = tmp_26_28_conv_naive_39 + computed_24_26_28_conv_naive_39
		}
		var computed_22_24_26_28_conv_naive_39 sint8 = kernel_conv_29[24] * conv_in[812]
		tmp_26_28_conv_naive_39 = tmp_26_28_conv_naive_39 + computed_22_24_26_28_conv_naive_39
		out_conv_naive_39[168] = tmp_26_28_conv_naive_39
		var res_conv_29 []sint8 = out_conv_naive_39
		var out_start_conv_29 int = i_conv_29 * 169
		copy(output_conv_29[out_start_conv_29:], res_conv_29)
	}
	var kernel_20_conv_29 []sint8 = kernel[100:125]
	var out_conv_naive_40 []sint8 = make([]sint8, 169)
	var tmp_25_27_conv_naive_40 sint8 = 0
	var computed_21_23_25_27_conv_naive_40 sint8 = kernel_20_conv_29[0] * conv_in[0]
	tmp_25_27_conv_naive_40 = tmp_25_27_conv_naive_40 + computed_21_23_25_27_conv_naive_40
	for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
		var convPos_23_25_27_conv_naive_40 int = wx_conv_naive_40 + 0
		var computed_23_25_27_conv_naive_40 sint8 = kernel_20_conv_29[convPos_23_25_27_conv_naive_40] * conv_in[0+(0+wx_conv_naive_40)]
		tmp_25_27_conv_naive_40 = tmp_25_27_conv_naive_40 + computed_23_25_27_conv_naive_40
	}
	var computed_22_23_25_27_conv_naive_40 sint8 = kernel_20_conv_29[4] * conv_in[4]
	tmp_25_27_conv_naive_40 = tmp_25_27_conv_naive_40 + computed_22_23_25_27_conv_naive_40
	for wy_conv_naive_40 := 1; wy_conv_naive_40 < 4; wy_conv_naive_40 = wy_conv_naive_40 + 1 {
		var convPos_21_25_27_conv_naive_40 int = 0 + wy_conv_naive_40*5
		var computed_21_25_27_conv_naive_40 sint8 = kernel_20_conv_29[convPos_21_25_27_conv_naive_40] * conv_in[(0+wy_conv_naive_40)*28+(0)]
		tmp_25_27_conv_naive_40 = tmp_25_27_conv_naive_40 + computed_21_25_27_conv_naive_40
		for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
			var convPos_25_27_conv_naive_40 int = wx_conv_naive_40 + wy_conv_naive_40*5
			var computed_25_27_conv_naive_40 sint8 = kernel_20_conv_29[convPos_25_27_conv_naive_40] * conv_in[(0+wy_conv_naive_40)*28+(0+wx_conv_naive_40)]
			tmp_25_27_conv_naive_40 = tmp_25_27_conv_naive_40 + computed_25_27_conv_naive_40
		}
		var convPos_22_25_27_conv_naive_40 int = 4 + wy_conv_naive_40*5
		var computed_22_25_27_conv_naive_40 sint8 = kernel_20_conv_29[convPos_22_25_27_conv_naive_40] * conv_in[(0+wy_conv_naive_40)*28+(4)]
		tmp_25_27_conv_naive_40 = tmp_25_27_conv_naive_40 + computed_22_25_27_conv_naive_40
	}
	var computed_21_24_25_27_conv_naive_40 sint8 = kernel_20_conv_29[20] * conv_in[112]
	tmp_25_27_conv_naive_40 = tmp_25_27_conv_naive_40 + computed_21_24_25_27_conv_naive_40
	for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
		var convPos_24_25_27_conv_naive_40 int = wx_conv_naive_40 + 20
		var computed_24_25_27_conv_naive_40 sint8 = kernel_20_conv_29[convPos_24_25_27_conv_naive_40] * conv_in[112+(0+wx_conv_naive_40)]
		tmp_25_27_conv_naive_40 = tmp_25_27_conv_naive_40 + computed_24_25_27_conv_naive_40
	}
	var computed_22_24_25_27_conv_naive_40 sint8 = kernel_20_conv_29[24] * conv_in[116]
	tmp_25_27_conv_naive_40 = tmp_25_27_conv_naive_40 + computed_22_24_25_27_conv_naive_40
	out_conv_naive_40[0] = tmp_25_27_conv_naive_40
	for x_conv_naive_40 := 1; x_conv_naive_40 < 12; x_conv_naive_40 = x_conv_naive_40 + 1 {
		var oPos_27_conv_naive_40 int = x_conv_naive_40 + 0
		var tmp_27_conv_naive_40 sint8 = 0
		var computed_21_23_27_conv_naive_40 sint8 = kernel_20_conv_29[0] * conv_in[0+(x_conv_naive_40*2+0)]
		tmp_27_conv_naive_40 = tmp_27_conv_naive_40 + computed_21_23_27_conv_naive_40
		for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
			var convPos_23_27_conv_naive_40 int = wx_conv_naive_40 + 0
			var computed_23_27_conv_naive_40 sint8 = kernel_20_conv_29[convPos_23_27_conv_naive_40] * conv_in[0+(x_conv_naive_40*2+wx_conv_naive_40)]
			tmp_27_conv_naive_40 = tmp_27_conv_naive_40 + computed_23_27_conv_naive_40
		}
		var computed_22_23_27_conv_naive_40 sint8 = kernel_20_conv_29[4] * conv_in[0+(x_conv_naive_40*2+4)]
		tmp_27_conv_naive_40 = tmp_27_conv_naive_40 + computed_22_23_27_conv_naive_40
		for wy_conv_naive_40 := 1; wy_conv_naive_40 < 4; wy_conv_naive_40 = wy_conv_naive_40 + 1 {
			var convPos_21_27_conv_naive_40 int = 0 + wy_conv_naive_40*5
			var computed_21_27_conv_naive_40 sint8 = kernel_20_conv_29[convPos_21_27_conv_naive_40] * conv_in[(0+wy_conv_naive_40)*28+(x_conv_naive_40*2+0)]
			tmp_27_conv_naive_40 = tmp_27_conv_naive_40 + computed_21_27_conv_naive_40
			for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
				var convPos_27_conv_naive_40 int = wx_conv_naive_40 + wy_conv_naive_40*5
				var computed_27_conv_naive_40 sint8 = kernel_20_conv_29[convPos_27_conv_naive_40] * conv_in[(0+wy_conv_naive_40)*28+(x_conv_naive_40*2+wx_conv_naive_40)]
				tmp_27_conv_naive_40 = tmp_27_conv_naive_40 + computed_27_conv_naive_40
			}
			var convPos_22_27_conv_naive_40 int = 4 + wy_conv_naive_40*5
			var computed_22_27_conv_naive_40 sint8 = kernel_20_conv_29[convPos_22_27_conv_naive_40] * conv_in[(0+wy_conv_naive_40)*28+(x_conv_naive_40*2+4)]
			tmp_27_conv_naive_40 = tmp_27_conv_naive_40 + computed_22_27_conv_naive_40
		}
		var computed_21_24_27_conv_naive_40 sint8 = kernel_20_conv_29[20] * conv_in[112+(x_conv_naive_40*2+0)]
		tmp_27_conv_naive_40 = tmp_27_conv_naive_40 + computed_21_24_27_conv_naive_40
		for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
			var convPos_24_27_conv_naive_40 int = wx_conv_naive_40 + 20
			var computed_24_27_conv_naive_40 sint8 = kernel_20_conv_29[convPos_24_27_conv_naive_40] * conv_in[112+(x_conv_naive_40*2+wx_conv_naive_40)]
			tmp_27_conv_naive_40 = tmp_27_conv_naive_40 + computed_24_27_conv_naive_40
		}
		var computed_22_24_27_conv_naive_40 sint8 = kernel_20_conv_29[24] * conv_in[112+(x_conv_naive_40*2+4)]
		tmp_27_conv_naive_40 = tmp_27_conv_naive_40 + computed_22_24_27_conv_naive_40
		out_conv_naive_40[oPos_27_conv_naive_40] = tmp_27_conv_naive_40
	}
	var tmp_26_27_conv_naive_40 sint8 = 0
	var computed_21_23_26_27_conv_naive_40 sint8 = kernel_20_conv_29[0] * conv_in[24]
	tmp_26_27_conv_naive_40 = tmp_26_27_conv_naive_40 + computed_21_23_26_27_conv_naive_40
	for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
		var convPos_23_26_27_conv_naive_40 int = wx_conv_naive_40 + 0
		var computed_23_26_27_conv_naive_40 sint8 = kernel_20_conv_29[convPos_23_26_27_conv_naive_40] * conv_in[0+(24+wx_conv_naive_40)]
		tmp_26_27_conv_naive_40 = tmp_26_27_conv_naive_40 + computed_23_26_27_conv_naive_40
	}
	var computed_22_23_26_27_conv_naive_40 sint8 = kernel_20_conv_29[4] * conv_in[28]
	tmp_26_27_conv_naive_40 = tmp_26_27_conv_naive_40 + computed_22_23_26_27_conv_naive_40
	for wy_conv_naive_40 := 1; wy_conv_naive_40 < 4; wy_conv_naive_40 = wy_conv_naive_40 + 1 {
		var convPos_21_26_27_conv_naive_40 int = 0 + wy_conv_naive_40*5
		var computed_21_26_27_conv_naive_40 sint8 = kernel_20_conv_29[convPos_21_26_27_conv_naive_40] * conv_in[(0+wy_conv_naive_40)*28+(24)]
		tmp_26_27_conv_naive_40 = tmp_26_27_conv_naive_40 + computed_21_26_27_conv_naive_40
		for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
			var convPos_26_27_conv_naive_40 int = wx_conv_naive_40 + wy_conv_naive_40*5
			var computed_26_27_conv_naive_40 sint8 = kernel_20_conv_29[convPos_26_27_conv_naive_40] * conv_in[(0+wy_conv_naive_40)*28+(24+wx_conv_naive_40)]
			tmp_26_27_conv_naive_40 = tmp_26_27_conv_naive_40 + computed_26_27_conv_naive_40
		}
		var convPos_22_26_27_conv_naive_40 int = 4 + wy_conv_naive_40*5
		var computed_22_26_27_conv_naive_40 sint8 = kernel_20_conv_29[convPos_22_26_27_conv_naive_40] * conv_in[(0+wy_conv_naive_40)*28+(28)]
		tmp_26_27_conv_naive_40 = tmp_26_27_conv_naive_40 + computed_22_26_27_conv_naive_40
	}
	var computed_21_24_26_27_conv_naive_40 sint8 = kernel_20_conv_29[20] * conv_in[136]
	tmp_26_27_conv_naive_40 = tmp_26_27_conv_naive_40 + computed_21_24_26_27_conv_naive_40
	for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
		var convPos_24_26_27_conv_naive_40 int = wx_conv_naive_40 + 20
		var computed_24_26_27_conv_naive_40 sint8 = kernel_20_conv_29[convPos_24_26_27_conv_naive_40] * conv_in[112+(24+wx_conv_naive_40)]
		tmp_26_27_conv_naive_40 = tmp_26_27_conv_naive_40 + computed_24_26_27_conv_naive_40
	}
	var computed_22_24_26_27_conv_naive_40 sint8 = kernel_20_conv_29[24] * conv_in[140]
	tmp_26_27_conv_naive_40 = tmp_26_27_conv_naive_40 + computed_22_24_26_27_conv_naive_40
	out_conv_naive_40[12] = tmp_26_27_conv_naive_40
	for y_conv_naive_40 := 1; y_conv_naive_40 < 12; y_conv_naive_40 = y_conv_naive_40 + 1 {
		var oPos_25_conv_naive_40 int = 0 + y_conv_naive_40*13
		var tmp_25_conv_naive_40 sint8 = 0
		var computed_21_23_25_conv_naive_40 sint8 = kernel_20_conv_29[0] * conv_in[(y_conv_naive_40*2+0)*28+(0)]
		tmp_25_conv_naive_40 = tmp_25_conv_naive_40 + computed_21_23_25_conv_naive_40
		for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
			var convPos_23_25_conv_naive_40 int = wx_conv_naive_40 + 0
			var computed_23_25_conv_naive_40 sint8 = kernel_20_conv_29[convPos_23_25_conv_naive_40] * conv_in[(y_conv_naive_40*2+0)*28+(0+wx_conv_naive_40)]
			tmp_25_conv_naive_40 = tmp_25_conv_naive_40 + computed_23_25_conv_naive_40
		}
		var computed_22_23_25_conv_naive_40 sint8 = kernel_20_conv_29[4] * conv_in[(y_conv_naive_40*2+0)*28+(4)]
		tmp_25_conv_naive_40 = tmp_25_conv_naive_40 + computed_22_23_25_conv_naive_40
		for wy_conv_naive_40 := 1; wy_conv_naive_40 < 4; wy_conv_naive_40 = wy_conv_naive_40 + 1 {
			var convPos_21_25_conv_naive_40 int = 0 + wy_conv_naive_40*5
			var computed_21_25_conv_naive_40 sint8 = kernel_20_conv_29[convPos_21_25_conv_naive_40] * conv_in[(y_conv_naive_40*2+wy_conv_naive_40)*28+(0)]
			tmp_25_conv_naive_40 = tmp_25_conv_naive_40 + computed_21_25_conv_naive_40
			for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
				var convPos_25_conv_naive_40 int = wx_conv_naive_40 + wy_conv_naive_40*5
				var computed_25_conv_naive_40 sint8 = kernel_20_conv_29[convPos_25_conv_naive_40] * conv_in[(y_conv_naive_40*2+wy_conv_naive_40)*28+(0+wx_conv_naive_40)]
				tmp_25_conv_naive_40 = tmp_25_conv_naive_40 + computed_25_conv_naive_40
			}
			var convPos_22_25_conv_naive_40 int = 4 + wy_conv_naive_40*5
			var computed_22_25_conv_naive_40 sint8 = kernel_20_conv_29[convPos_22_25_conv_naive_40] * conv_in[(y_conv_naive_40*2+wy_conv_naive_40)*28+(4)]
			tmp_25_conv_naive_40 = tmp_25_conv_naive_40 + computed_22_25_conv_naive_40
		}
		var computed_21_24_25_conv_naive_40 sint8 = kernel_20_conv_29[20] * conv_in[(y_conv_naive_40*2+4)*28+(0)]
		tmp_25_conv_naive_40 = tmp_25_conv_naive_40 + computed_21_24_25_conv_naive_40
		for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
			var convPos_24_25_conv_naive_40 int = wx_conv_naive_40 + 20
			var computed_24_25_conv_naive_40 sint8 = kernel_20_conv_29[convPos_24_25_conv_naive_40] * conv_in[(y_conv_naive_40*2+4)*28+(0+wx_conv_naive_40)]
			tmp_25_conv_naive_40 = tmp_25_conv_naive_40 + computed_24_25_conv_naive_40
		}
		var computed_22_24_25_conv_naive_40 sint8 = kernel_20_conv_29[24] * conv_in[(y_conv_naive_40*2+4)*28+(4)]
		tmp_25_conv_naive_40 = tmp_25_conv_naive_40 + computed_22_24_25_conv_naive_40
		out_conv_naive_40[oPos_25_conv_naive_40] = tmp_25_conv_naive_40
		for x_conv_naive_40 := 1; x_conv_naive_40 < 12; x_conv_naive_40 = x_conv_naive_40 + 1 {
			var oPos_conv_naive_40 int = x_conv_naive_40 + y_conv_naive_40*13
			var tmp_conv_naive_40 sint8 = 0
			var computed_21_23_conv_naive_40 sint8 = kernel_20_conv_29[0] * conv_in[(y_conv_naive_40*2+0)*28+(x_conv_naive_40*2+0)]
			tmp_conv_naive_40 = tmp_conv_naive_40 + computed_21_23_conv_naive_40
			for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
				var convPos_23_conv_naive_40 int = wx_conv_naive_40 + 0
				var computed_23_conv_naive_40 sint8 = kernel_20_conv_29[convPos_23_conv_naive_40] * conv_in[(y_conv_naive_40*2+0)*28+(x_conv_naive_40*2+wx_conv_naive_40)]
				tmp_conv_naive_40 = tmp_conv_naive_40 + computed_23_conv_naive_40
			}
			var computed_22_23_conv_naive_40 sint8 = kernel_20_conv_29[4] * conv_in[(y_conv_naive_40*2+0)*28+(x_conv_naive_40*2+4)]
			tmp_conv_naive_40 = tmp_conv_naive_40 + computed_22_23_conv_naive_40
			for wy_conv_naive_40 := 1; wy_conv_naive_40 < 4; wy_conv_naive_40 = wy_conv_naive_40 + 1 {
				var convPos_21_conv_naive_40 int = 0 + wy_conv_naive_40*5
				var computed_21_conv_naive_40 sint8 = kernel_20_conv_29[convPos_21_conv_naive_40] * conv_in[(y_conv_naive_40*2+wy_conv_naive_40)*28+(x_conv_naive_40*2+0)]
				tmp_conv_naive_40 = tmp_conv_naive_40 + computed_21_conv_naive_40
				for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
					var convPos_conv_naive_40 int = wx_conv_naive_40 + wy_conv_naive_40*5
					var computed_conv_naive_40 sint8 = kernel_20_conv_29[convPos_conv_naive_40] * conv_in[(y_conv_naive_40*2+wy_conv_naive_40)*28+(x_conv_naive_40*2+wx_conv_naive_40)]
					tmp_conv_naive_40 = tmp_conv_naive_40 + computed_conv_naive_40
				}
				var convPos_22_conv_naive_40 int = 4 + wy_conv_naive_40*5
				var computed_22_conv_naive_40 sint8 = kernel_20_conv_29[convPos_22_conv_naive_40] * conv_in[(y_conv_naive_40*2+wy_conv_naive_40)*28+(x_conv_naive_40*2+4)]
				tmp_conv_naive_40 = tmp_conv_naive_40 + computed_22_conv_naive_40
			}
			var computed_21_24_conv_naive_40 sint8 = kernel_20_conv_29[20] * conv_in[(y_conv_naive_40*2+4)*28+(x_conv_naive_40*2+0)]
			tmp_conv_naive_40 = tmp_conv_naive_40 + computed_21_24_conv_naive_40
			for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
				var convPos_24_conv_naive_40 int = wx_conv_naive_40 + 20
				var computed_24_conv_naive_40 sint8 = kernel_20_conv_29[convPos_24_conv_naive_40] * conv_in[(y_conv_naive_40*2+4)*28+(x_conv_naive_40*2+wx_conv_naive_40)]
				tmp_conv_naive_40 = tmp_conv_naive_40 + computed_24_conv_naive_40
			}
			var computed_22_24_conv_naive_40 sint8 = kernel_20_conv_29[24] * conv_in[(y_conv_naive_40*2+4)*28+(x_conv_naive_40*2+4)]
			tmp_conv_naive_40 = tmp_conv_naive_40 + computed_22_24_conv_naive_40
			out_conv_naive_40[oPos_conv_naive_40] = tmp_conv_naive_40
		}
		var oPos_26_conv_naive_40 int = 12 + y_conv_naive_40*13
		var tmp_26_conv_naive_40 sint8 = 0
		var computed_21_23_26_conv_naive_40 sint8 = kernel_20_conv_29[0] * conv_in[(y_conv_naive_40*2+0)*28+(24)]
		tmp_26_conv_naive_40 = tmp_26_conv_naive_40 + computed_21_23_26_conv_naive_40
		for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
			var convPos_23_26_conv_naive_40 int = wx_conv_naive_40 + 0
			var computed_23_26_conv_naive_40 sint8 = kernel_20_conv_29[convPos_23_26_conv_naive_40] * conv_in[(y_conv_naive_40*2+0)*28+(24+wx_conv_naive_40)]
			tmp_26_conv_naive_40 = tmp_26_conv_naive_40 + computed_23_26_conv_naive_40
		}
		var computed_22_23_26_conv_naive_40 sint8 = kernel_20_conv_29[4] * conv_in[(y_conv_naive_40*2+0)*28+(28)]
		tmp_26_conv_naive_40 = tmp_26_conv_naive_40 + computed_22_23_26_conv_naive_40
		for wy_conv_naive_40 := 1; wy_conv_naive_40 < 4; wy_conv_naive_40 = wy_conv_naive_40 + 1 {
			var convPos_21_26_conv_naive_40 int = 0 + wy_conv_naive_40*5
			var computed_21_26_conv_naive_40 sint8 = kernel_20_conv_29[convPos_21_26_conv_naive_40] * conv_in[(y_conv_naive_40*2+wy_conv_naive_40)*28+(24)]
			tmp_26_conv_naive_40 = tmp_26_conv_naive_40 + computed_21_26_conv_naive_40
			for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
				var convPos_26_conv_naive_40 int = wx_conv_naive_40 + wy_conv_naive_40*5
				var computed_26_conv_naive_40 sint8 = kernel_20_conv_29[convPos_26_conv_naive_40] * conv_in[(y_conv_naive_40*2+wy_conv_naive_40)*28+(24+wx_conv_naive_40)]
				tmp_26_conv_naive_40 = tmp_26_conv_naive_40 + computed_26_conv_naive_40
			}
			var convPos_22_26_conv_naive_40 int = 4 + wy_conv_naive_40*5
			var computed_22_26_conv_naive_40 sint8 = kernel_20_conv_29[convPos_22_26_conv_naive_40] * conv_in[(y_conv_naive_40*2+wy_conv_naive_40)*28+(28)]
			tmp_26_conv_naive_40 = tmp_26_conv_naive_40 + computed_22_26_conv_naive_40
		}
		var computed_21_24_26_conv_naive_40 sint8 = kernel_20_conv_29[20] * conv_in[(y_conv_naive_40*2+4)*28+(24)]
		tmp_26_conv_naive_40 = tmp_26_conv_naive_40 + computed_21_24_26_conv_naive_40
		for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
			var convPos_24_26_conv_naive_40 int = wx_conv_naive_40 + 20
			var computed_24_26_conv_naive_40 sint8 = kernel_20_conv_29[convPos_24_26_conv_naive_40] * conv_in[(y_conv_naive_40*2+4)*28+(24+wx_conv_naive_40)]
			tmp_26_conv_naive_40 = tmp_26_conv_naive_40 + computed_24_26_conv_naive_40
		}
		var computed_22_24_26_conv_naive_40 sint8 = kernel_20_conv_29[24] * conv_in[(y_conv_naive_40*2+4)*28+(28)]
		tmp_26_conv_naive_40 = tmp_26_conv_naive_40 + computed_22_24_26_conv_naive_40
		out_conv_naive_40[oPos_26_conv_naive_40] = tmp_26_conv_naive_40
	}
	var tmp_25_28_conv_naive_40 sint8 = 0
	var computed_21_23_25_28_conv_naive_40 sint8 = kernel_20_conv_29[0] * conv_in[672]
	tmp_25_28_conv_naive_40 = tmp_25_28_conv_naive_40 + computed_21_23_25_28_conv_naive_40
	for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
		var convPos_23_25_28_conv_naive_40 int = wx_conv_naive_40 + 0
		var computed_23_25_28_conv_naive_40 sint8 = kernel_20_conv_29[convPos_23_25_28_conv_naive_40] * conv_in[672+(0+wx_conv_naive_40)]
		tmp_25_28_conv_naive_40 = tmp_25_28_conv_naive_40 + computed_23_25_28_conv_naive_40
	}
	var computed_22_23_25_28_conv_naive_40 sint8 = kernel_20_conv_29[4] * conv_in[676]
	tmp_25_28_conv_naive_40 = tmp_25_28_conv_naive_40 + computed_22_23_25_28_conv_naive_40
	for wy_conv_naive_40 := 1; wy_conv_naive_40 < 4; wy_conv_naive_40 = wy_conv_naive_40 + 1 {
		var convPos_21_25_28_conv_naive_40 int = 0 + wy_conv_naive_40*5
		var computed_21_25_28_conv_naive_40 sint8 = kernel_20_conv_29[convPos_21_25_28_conv_naive_40] * conv_in[(24+wy_conv_naive_40)*28+(0)]
		tmp_25_28_conv_naive_40 = tmp_25_28_conv_naive_40 + computed_21_25_28_conv_naive_40
		for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
			var convPos_25_28_conv_naive_40 int = wx_conv_naive_40 + wy_conv_naive_40*5
			var computed_25_28_conv_naive_40 sint8 = kernel_20_conv_29[convPos_25_28_conv_naive_40] * conv_in[(24+wy_conv_naive_40)*28+(0+wx_conv_naive_40)]
			tmp_25_28_conv_naive_40 = tmp_25_28_conv_naive_40 + computed_25_28_conv_naive_40
		}
		var convPos_22_25_28_conv_naive_40 int = 4 + wy_conv_naive_40*5
		var computed_22_25_28_conv_naive_40 sint8 = kernel_20_conv_29[convPos_22_25_28_conv_naive_40] * conv_in[(24+wy_conv_naive_40)*28+(4)]
		tmp_25_28_conv_naive_40 = tmp_25_28_conv_naive_40 + computed_22_25_28_conv_naive_40
	}
	var computed_21_24_25_28_conv_naive_40 sint8 = kernel_20_conv_29[20] * conv_in[784]
	tmp_25_28_conv_naive_40 = tmp_25_28_conv_naive_40 + computed_21_24_25_28_conv_naive_40
	for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
		var convPos_24_25_28_conv_naive_40 int = wx_conv_naive_40 + 20
		var computed_24_25_28_conv_naive_40 sint8 = kernel_20_conv_29[convPos_24_25_28_conv_naive_40] * conv_in[784+(0+wx_conv_naive_40)]
		tmp_25_28_conv_naive_40 = tmp_25_28_conv_naive_40 + computed_24_25_28_conv_naive_40
	}
	var computed_22_24_25_28_conv_naive_40 sint8 = kernel_20_conv_29[24] * conv_in[788]
	tmp_25_28_conv_naive_40 = tmp_25_28_conv_naive_40 + computed_22_24_25_28_conv_naive_40
	out_conv_naive_40[156] = tmp_25_28_conv_naive_40
	for x_conv_naive_40 := 1; x_conv_naive_40 < 12; x_conv_naive_40 = x_conv_naive_40 + 1 {
		var oPos_28_conv_naive_40 int = x_conv_naive_40 + 156
		var tmp_28_conv_naive_40 sint8 = 0
		var computed_21_23_28_conv_naive_40 sint8 = kernel_20_conv_29[0] * conv_in[672+(x_conv_naive_40*2+0)]
		tmp_28_conv_naive_40 = tmp_28_conv_naive_40 + computed_21_23_28_conv_naive_40
		for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
			var convPos_23_28_conv_naive_40 int = wx_conv_naive_40 + 0
			var computed_23_28_conv_naive_40 sint8 = kernel_20_conv_29[convPos_23_28_conv_naive_40] * conv_in[672+(x_conv_naive_40*2+wx_conv_naive_40)]
			tmp_28_conv_naive_40 = tmp_28_conv_naive_40 + computed_23_28_conv_naive_40
		}
		var computed_22_23_28_conv_naive_40 sint8 = kernel_20_conv_29[4] * conv_in[672+(x_conv_naive_40*2+4)]
		tmp_28_conv_naive_40 = tmp_28_conv_naive_40 + computed_22_23_28_conv_naive_40
		for wy_conv_naive_40 := 1; wy_conv_naive_40 < 4; wy_conv_naive_40 = wy_conv_naive_40 + 1 {
			var convPos_21_28_conv_naive_40 int = 0 + wy_conv_naive_40*5
			var computed_21_28_conv_naive_40 sint8 = kernel_20_conv_29[convPos_21_28_conv_naive_40] * conv_in[(24+wy_conv_naive_40)*28+(x_conv_naive_40*2+0)]
			tmp_28_conv_naive_40 = tmp_28_conv_naive_40 + computed_21_28_conv_naive_40
			for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
				var convPos_28_conv_naive_40 int = wx_conv_naive_40 + wy_conv_naive_40*5
				var computed_28_conv_naive_40 sint8 = kernel_20_conv_29[convPos_28_conv_naive_40] * conv_in[(24+wy_conv_naive_40)*28+(x_conv_naive_40*2+wx_conv_naive_40)]
				tmp_28_conv_naive_40 = tmp_28_conv_naive_40 + computed_28_conv_naive_40
			}
			var convPos_22_28_conv_naive_40 int = 4 + wy_conv_naive_40*5
			var computed_22_28_conv_naive_40 sint8 = kernel_20_conv_29[convPos_22_28_conv_naive_40] * conv_in[(24+wy_conv_naive_40)*28+(x_conv_naive_40*2+4)]
			tmp_28_conv_naive_40 = tmp_28_conv_naive_40 + computed_22_28_conv_naive_40
		}
		var computed_21_24_28_conv_naive_40 sint8 = kernel_20_conv_29[20] * conv_in[784+(x_conv_naive_40*2+0)]
		tmp_28_conv_naive_40 = tmp_28_conv_naive_40 + computed_21_24_28_conv_naive_40
		for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
			var convPos_24_28_conv_naive_40 int = wx_conv_naive_40 + 20
			var computed_24_28_conv_naive_40 sint8 = kernel_20_conv_29[convPos_24_28_conv_naive_40] * conv_in[784+(x_conv_naive_40*2+wx_conv_naive_40)]
			tmp_28_conv_naive_40 = tmp_28_conv_naive_40 + computed_24_28_conv_naive_40
		}
		var computed_22_24_28_conv_naive_40 sint8 = kernel_20_conv_29[24] * conv_in[784+(x_conv_naive_40*2+4)]
		tmp_28_conv_naive_40 = tmp_28_conv_naive_40 + computed_22_24_28_conv_naive_40
		out_conv_naive_40[oPos_28_conv_naive_40] = tmp_28_conv_naive_40
	}
	var tmp_26_28_conv_naive_40 sint8 = 0
	var computed_21_23_26_28_conv_naive_40 sint8 = kernel_20_conv_29[0] * conv_in[696]
	tmp_26_28_conv_naive_40 = tmp_26_28_conv_naive_40 + computed_21_23_26_28_conv_naive_40
	for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
		var convPos_23_26_28_conv_naive_40 int = wx_conv_naive_40 + 0
		var computed_23_26_28_conv_naive_40 sint8 = kernel_20_conv_29[convPos_23_26_28_conv_naive_40] * conv_in[672+(24+wx_conv_naive_40)]
		tmp_26_28_conv_naive_40 = tmp_26_28_conv_naive_40 + computed_23_26_28_conv_naive_40
	}
	var computed_22_23_26_28_conv_naive_40 sint8 = kernel_20_conv_29[4] * conv_in[700]
	tmp_26_28_conv_naive_40 = tmp_26_28_conv_naive_40 + computed_22_23_26_28_conv_naive_40
	for wy_conv_naive_40 := 1; wy_conv_naive_40 < 4; wy_conv_naive_40 = wy_conv_naive_40 + 1 {
		var convPos_21_26_28_conv_naive_40 int = 0 + wy_conv_naive_40*5
		var computed_21_26_28_conv_naive_40 sint8 = kernel_20_conv_29[convPos_21_26_28_conv_naive_40] * conv_in[(24+wy_conv_naive_40)*28+(24)]
		tmp_26_28_conv_naive_40 = tmp_26_28_conv_naive_40 + computed_21_26_28_conv_naive_40
		for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
			var convPos_26_28_conv_naive_40 int = wx_conv_naive_40 + wy_conv_naive_40*5
			var computed_26_28_conv_naive_40 sint8 = kernel_20_conv_29[convPos_26_28_conv_naive_40] * conv_in[(24+wy_conv_naive_40)*28+(24+wx_conv_naive_40)]
			tmp_26_28_conv_naive_40 = tmp_26_28_conv_naive_40 + computed_26_28_conv_naive_40
		}
		var convPos_22_26_28_conv_naive_40 int = 4 + wy_conv_naive_40*5
		var computed_22_26_28_conv_naive_40 sint8 = kernel_20_conv_29[convPos_22_26_28_conv_naive_40] * conv_in[(24+wy_conv_naive_40)*28+(28)]
		tmp_26_28_conv_naive_40 = tmp_26_28_conv_naive_40 + computed_22_26_28_conv_naive_40
	}
	var computed_21_24_26_28_conv_naive_40 sint8 = kernel_20_conv_29[20] * conv_in[808]
	tmp_26_28_conv_naive_40 = tmp_26_28_conv_naive_40 + computed_21_24_26_28_conv_naive_40
	for wx_conv_naive_40 := 1; wx_conv_naive_40 < 4; wx_conv_naive_40 = wx_conv_naive_40 + 1 {
		var convPos_24_26_28_conv_naive_40 int = wx_conv_naive_40 + 20
		var computed_24_26_28_conv_naive_40 sint8 = kernel_20_conv_29[convPos_24_26_28_conv_naive_40] * conv_in[784+(24+wx_conv_naive_40)]
		tmp_26_28_conv_naive_40 = tmp_26_28_conv_naive_40 + computed_24_26_28_conv_naive_40
	}
	var computed_22_24_26_28_conv_naive_40 sint8 = kernel_20_conv_29[24] * conv_in[812]
	tmp_26_28_conv_naive_40 = tmp_26_28_conv_naive_40 + computed_22_24_26_28_conv_naive_40
	out_conv_naive_40[168] = tmp_26_28_conv_naive_40
	var res_20_conv_29 []sint8 = out_conv_naive_40
	copy(output_conv_29[676:], res_20_conv_29)
	var conv_out []sint8 = output_conv_29
	var val_7 sint8 = conv_out[0]
	var ret_sqr_30 sint8 = val_7 * val_7
	conv_out[0] = ret_sqr_30
	for i := 1; i < 844; i = i + 1 {
		var val sint8 = conv_out[i]
		var ret_sqr_31 sint8 = val * val
		conv_out[i] = ret_sqr_31
	}
	var val_8 sint8 = conv_out[844]
	var ret_sqr_32 sint8 = val_8 * val_8
	conv_out[844] = ret_sqr_32
	var out_mmul0_33 []sint8 = make([]sint8, 100)
	var aRow_13_mmul0_33 []sint8 = pool[0:845]
	var sum_13_mmul0_33 sint8 = 0
	var tmp_11_13_mmul0_33 = aRow_13_mmul0_33[0] * conv_out[0]
	sum_13_mmul0_33 = sum_13_mmul0_33 + tmp_11_13_mmul0_33
	for j_mmul0_33 := 1; j_mmul0_33 < 844; j_mmul0_33 = j_mmul0_33 + 1 {
		var tmp_13_mmul0_33 = aRow_13_mmul0_33[j_mmul0_33] * conv_out[j_mmul0_33]
		sum_13_mmul0_33 = sum_13_mmul0_33 + tmp_13_mmul0_33
	}
	var tmp_12_13_mmul0_33 = aRow_13_mmul0_33[844] * conv_out[844]
	sum_13_mmul0_33 = sum_13_mmul0_33 + tmp_12_13_mmul0_33
	out_mmul0_33[0] = sum_13_mmul0_33
	for i_mmul0_33 := 1; i_mmul0_33 < 99; i_mmul0_33 = i_mmul0_33 + 1 {
		var aRow_mmul0_33 []sint8 = pool[i_mmul0_33*845 : (i_mmul0_33+1)*845]
		var sum_mmul0_33 sint8 = 0
		var tmp_11_mmul0_33 = aRow_mmul0_33[0] * conv_out[0]
		sum_mmul0_33 = sum_mmul0_33 + tmp_11_mmul0_33
		for j_mmul0_33 := 1; j_mmul0_33 < 844; j_mmul0_33 = j_mmul0_33 + 1 {
			var tmp_mmul0_33 = aRow_mmul0_33[j_mmul0_33] * conv_out[j_mmul0_33]
			sum_mmul0_33 = sum_mmul0_33 + tmp_mmul0_33
		}
		var tmp_12_mmul0_33 = aRow_mmul0_33[844] * conv_out[844]
		sum_mmul0_33 = sum_mmul0_33 + tmp_12_mmul0_33
		out_mmul0_33[i_mmul0_33] = sum_mmul0_33
	}
	var aRow_14_mmul0_33 []sint8 = pool[83655:84500]
	var sum_14_mmul0_33 sint8 = 0
	var tmp_11_14_mmul0_33 = aRow_14_mmul0_33[0] * conv_out[0]
	sum_14_mmul0_33 = sum_14_mmul0_33 + tmp_11_14_mmul0_33
	for j_mmul0_33 := 1; j_mmul0_33 < 844; j_mmul0_33 = j_mmul0_33 + 1 {
		var tmp_14_mmul0_33 = aRow_14_mmul0_33[j_mmul0_33] * conv_out[j_mmul0_33]
		sum_14_mmul0_33 = sum_14_mmul0_33 + tmp_14_mmul0_33
	}
	var tmp_12_14_mmul0_33 = aRow_14_mmul0_33[844] * conv_out[844]
	sum_14_mmul0_33 = sum_14_mmul0_33 + tmp_12_14_mmul0_33
	out_mmul0_33[99] = sum_14_mmul0_33
	var im_out []sint8 = out_mmul0_33
	var val__9 sint8 = im_out[0]
	var ret_sqr_34 sint8 = val__9 * val__9
	im_out[0] = ret_sqr_34
	for i := 1; i < 99; i = i + 1 {
		var val_ sint8 = im_out[i]
		var ret_sqr_35 sint8 = val_ * val_
		im_out[i] = ret_sqr_35
	}
	var val__10 sint8 = im_out[99]
	var ret_sqr_36 sint8 = val__10 * val__10
	im_out[99] = ret_sqr_36
	var out_mmul1_37 []sint8 = make([]sint8, 10)
	var aRow_17_mmul1_37 []sint8 = fc[0:100]
	var sum_17_mmul1_37 sint8 = 0
	var tmp_15_17_mmul1_37 sint8 = aRow_17_mmul1_37[0] * im_out[0]
	sum_17_mmul1_37 = sum_17_mmul1_37 + tmp_15_17_mmul1_37
	for j_mmul1_37 := 1; j_mmul1_37 < 99; j_mmul1_37 = j_mmul1_37 + 1 {
		var tmp_17_mmul1_37 sint8 = aRow_17_mmul1_37[j_mmul1_37] * im_out[j_mmul1_37]
		sum_17_mmul1_37 = sum_17_mmul1_37 + tmp_17_mmul1_37
	}
	var tmp_16_17_mmul1_37 sint8 = aRow_17_mmul1_37[99] * im_out[99]
	sum_17_mmul1_37 = sum_17_mmul1_37 + tmp_16_17_mmul1_37
	out_mmul1_37[0] = sum_17_mmul1_37
	for i_mmul1_37 := 1; i_mmul1_37 < 9; i_mmul1_37 = i_mmul1_37 + 1 {
		var aRow_mmul1_37 []sint8 = fc[i_mmul1_37*100 : (i_mmul1_37+1)*100]
		var sum_mmul1_37 sint8 = 0
		var tmp_15_mmul1_37 sint8 = aRow_mmul1_37[0] * im_out[0]
		sum_mmul1_37 = sum_mmul1_37 + tmp_15_mmul1_37
		for j_mmul1_37 := 1; j_mmul1_37 < 99; j_mmul1_37 = j_mmul1_37 + 1 {
			var tmp_mmul1_37 sint8 = aRow_mmul1_37[j_mmul1_37] * im_out[j_mmul1_37]
			sum_mmul1_37 = sum_mmul1_37 + tmp_mmul1_37
		}
		var tmp_16_mmul1_37 sint8 = aRow_mmul1_37[99] * im_out[99]
		sum_mmul1_37 = sum_mmul1_37 + tmp_16_mmul1_37
		out_mmul1_37[i_mmul1_37] = sum_mmul1_37
	}
	var aRow_18_mmul1_37 []sint8 = fc[900:1000]
	var sum_18_mmul1_37 sint8 = 0
	var tmp_15_18_mmul1_37 sint8 = aRow_18_mmul1_37[0] * im_out[0]
	sum_18_mmul1_37 = sum_18_mmul1_37 + tmp_15_18_mmul1_37
	for j_mmul1_37 := 1; j_mmul1_37 < 99; j_mmul1_37 = j_mmul1_37 + 1 {
		var tmp_18_mmul1_37 sint8 = aRow_18_mmul1_37[j_mmul1_37] * im_out[j_mmul1_37]
		sum_18_mmul1_37 = sum_18_mmul1_37 + tmp_18_mmul1_37
	}
	var tmp_16_18_mmul1_37 sint8 = aRow_18_mmul1_37[99] * im_out[99]
	sum_18_mmul1_37 = sum_18_mmul1_37 + tmp_16_18_mmul1_37
	out_mmul1_37[9] = sum_18_mmul1_37
	var final_out []sint8 = out_mmul1_37
	var out []int8 = openi8n(final_out, 10)
	print(out)
}
