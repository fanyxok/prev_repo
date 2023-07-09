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
	var image []sint8 = i8n(0, ImageWidth*ImageWidth)
	var kernel []sint8 = i8n(0, OutputChannels*WindowWidth*WindowWidth)
	var pool []sint8 = i8n(0, FcWidth*ConvSize*OutputChannels)
	var fc []sint8 = i8n(0, FinalOutputChannels*FcWidth)
	// padding
	var conv_in []sint8 = make([]sint8, PaddedWidth*PaddedWidth)
	for i := 0; i < PaddedWidth; i = i + 1 {
		conv_in[i] = 0
		conv_in[i+PaddedWidth] = 0
		conv_in[PaddedWidth*i] = 0
		conv_in[PaddedWidth*i+1] = 0
	}

	for y := 0; y < ImageWidth; y = y + 1 {
		for x := 0; x < ImageWidth; x = x + 1 {
			var img_pixel int = y*ImageWidth + x
			var in_pixel int = (y+2)*PaddedWidth + x + 2
			conv_in[in_pixel] = image[img_pixel]
		}
	}
	// Convolution (1) [OutputChannels * ConvSize]
	var conv_out []sint8 = conv(conv_in, kernel)

	// Activation Function (2)
	for i := 0; i < OutputChannels*ConvSize; i = i + 1 {
		var val sint8 = conv_out[i]
		conv_out[i] = sqr(val)
	}

	// Combination of Mean pooling and Fully connected (3)
	var im_out []sint8 = mmul0(pool, conv_out)
	// Activation Function (4)
	for i := 0; i < FcWidth; i = i + 1 {
		var val_ sint8 = im_out[i]
		im_out[i] = sqr(val_)
	}
	// Fully Connected (5)
	var final_out []sint8 = mmul1(fc, im_out)

	var out []int8 = openi8n(final_out, FinalOutputChannels)
	print(out)
}

// FULLY_CONNECTED_WIDTH, 1, OUTPUT_CHANNELS * SIZE_CONVOLUTION);
func mmul0(pool_layer []sint8, conv_layer []sint8) []sint8 {
	var out []sint8 = make([]sint8, FcWidth)
	var common int = OutputChannels * ConvSize
	for i := 0; i < FcWidth; i = i + 1 {
		var aRow []sint8 = pool_layer[i*common : (i+1)*common]
		var sum sint8 = 0
		for j := 0; j < OutputChannels*ConvSize; j = j + 1 {
			var tmp = aRow[j] * conv_layer[j]
			sum = sum + tmp
		}
		out[i] = sum
	}
	return out
}

// FINAL_OUTPUT_CHANNELS, 1, FULLY_CONNECTED_WIDTH);
func mmul1(fc []sint8, im_layer []sint8) []sint8 {
	var out []sint8 = make([]sint8, FinalOutputChannels)
	for i := 0; i < FinalOutputChannels; i = i + 1 {
		var aRow []sint8 = fc[i*FcWidth : (i+1)*FcWidth]
		var sum sint8 = 0
		for j := 0; j < FcWidth; j = j + 1 {
			var tmp sint8 = aRow[j] * im_layer[j]
			sum = sum + tmp
		}
		out[i] = sum
	}
	return out
}

func sqr(val sint8) sint8 {
	var ret sint8 = val * val
	return ret
}

func conv(image []sint8, kernels []sint8) []sint8 {
	var kernal_size int = WindowWidth * WindowWidth
	var output []sint8 = make([]sint8, OutputChannels*ConvSize)
	for i := 0; i < OutputChannels; i = i + 1 {
		var kernal_start int = i * kernal_size
		var tmp int = i + 1
		var kernal_end int = tmp * kernal_size
		var kernel []sint8 = kernels[kernal_start:kernal_end]
		var res []sint8 = conv_naive(image, kernel)
		var out_start int = i * ConvSize
		copy(output[out_start:], res)
	}
	return output
}

func conv_naive(image []sint8, kernel []sint8) []sint8 {
	var out []sint8 = make([]sint8, ConvSize)
	for y := 0; y < ConvWidth; y = y + 1 {
		for x := 0; x < ConvWidth; x = x + 1 {
			var oPos int = x + y*ConvWidth
			var tmp sint8 = 0
			for wy := 0; wy < WindowWidth; wy = wy + 1 {
				for wx := 0; wx < WindowWidth; wx = wx + 1 {
					var convPos int = wx + wy*WindowWidth
					var computed sint8 = kernel[convPos] * image[(y*Stride+wy)*ImageWidth+(x*Stride+wx)]
					tmp = tmp + computed
				}
			}
			out[oPos] = tmp
		}
	}
	return out
}
