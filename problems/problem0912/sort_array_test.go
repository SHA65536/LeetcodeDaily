package problem0912

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []int
	Expected []int
}

var Results = []Result{
	{[]int{5, 2, 3, 1}, []int{1, 2, 3, 5}},
	{[]int{5, 1, 1, 2, 0, 0}, []int{0, 0, 1, 1, 2, 5}},
	{[]int{2, 3, 1}, []int{1, 2, 3}},
	{[]int{2, 3, 1, 4, 5}, []int{1, 2, 3, 4, 5}},
	{
		[]int{349, 65, 537, 473, 917, 989, 103, 935, 475, 955, 15, 753, 286, 320, 737, 399, 958, 297, 893, 380, 663, 543, 28, 614, 384, 182, 682, 373, 229, 957, 438, 191, 388, 559, 729, 26, 389, 525, 88, 478, 453, 407, 622, 431, 89, 832, 91, 478, 689, 415, 812, 496, 531, 379, 186, 379, 738, 868, 315, 1, 91, 649, 978, 784, 95, 775, 250, 965, 780, 100, 412, 529, 29, 610, 394, 475, 634, 8, 533, 973, 800, 271, 786, 558, 88, 291, 443, 201, 912, 911, 27, 128, 636, 696, 272, 963, 446, 36, 865, 991, 329, 206, 669, 617, 335, 698, 842, 284, 128, 60, 638, 581, 419, 702, 365, 842, 791, 880, 103, 95, 357, 684, 95, 225, 716, 123, 417, 521, 312, 431, 278, 813, 678, 847, 952, 511, 763, 547, 26, 844, 242, 742, 124, 768, 501, 828, 328, 430, 342, 656, 451, 650, 398, 575, 373, 440, 945, 465, 454, 294, 147, 256, 813, 766, 533, 501, 97, 353, 520, 612, 4, 258, 242, 898, 850, 256, 844, 417, 768, 805, 940, 206, 25, 195, 146, 400, 466, 314, 668, 908, 407, 380, 690, 242, 221, 962, 313, 144, 568, 263, 364, 883, 897, 766, 379, 646, 692, 779, 900, 681, 942, 668, 223, 593, 656, 424, 441, 870, 891, 631, 958, 416, 448, 568, 71, 978, 931, 674, 96, 101, 171, 989, 789, 111, 702, 432, 818, 785, 482, 183, 347, 564, 105, 501, 483, 752, 218, 960, 695, 434, 360, 731, 304, 372, 542, 219, 785, 155, 194, 981, 155, 212, 197, 773, 94, 438, 595, 795, 534, 697, 736, 340, 363, 200, 264, 816, 48, 833, 501, 171, 195, 776, 396, 556, 874, 814, 647, 265, 253, 876, 929, 285, 402, 745, 26, 681, 798, 601, 781, 567, 538, 351, 189, 600, 694, 519, 600, 306, 148, 763, 706, 529, 651, 195, 821, 197, 826, 197, 257, 322, 809, 101, 495, 83, 829, 327, 491, 527, 261, 967, 106, 118, 839, 525, 494, 768, 411, 48, 650, 380, 678, 807, 519, 271, 417, 687, 118, 454, 217, 781, 696, 70, 684, 143, 765, 506, 935, 626, 416, 934, 111, 793, 243, 936, 123, 112, 975, 194, 735, 992, 858, 324, 514, 691, 536, 937, 180, 434, 36, 851, 209, 411, 1000, 424, 834, 864, 574, 689, 194, 76, 616, 613, 865, 961, 957, 145, 357, 946, 952, 143, 550, 234, 683, 915, 325, 398, 602, 526, 220, 90, 714, 803, 422, 656, 8, 526, 961, 487, 169, 813, 728, 45, 189, 375, 558, 764, 488, 474, 309, 458, 900, 92, 53, 695, 626, 163, 375, 606, 523, 695, 640, 45, 460, 805, 205, 300, 453, 398, 115, 178, 677, 682, 549, 452, 403, 421, 103, 660, 266, 889, 353, 114, 574, 126, 887, 410, 631, 662, 645, 187, 883, 319, 310, 229, 693, 243, 27, 472, 227, 965, 363, 474, 1000, 435, 413, 808, 66, 887, 976, 585, 436, 864, 663, 113, 703, 179, 336, 890, 741, 298, 662, 224, 362, 924, 210, 343, 261, 518, 733, 263, 983, 748, 515, 298, 565, 928, 709, 188, 325, 661, 483, 926, 542, 980, 304, 762, 286, 476, 493, 859, 523, 314, 315, 241, 376, 946, 767, 664, 597, 287, 445, 991, 841, 304, 161, 663, 56, 393, 231, 778, 788, 348, 761, 843, 531, 437, 405, 520, 749, 944, 608, 817, 991, 325, 589, 294, 994, 538, 580, 149, 740, 367, 856, 561, 980, 310, 237, 944, 935, 498, 856, 968, 283, 868, 61, 727, 664, 891, 905, 835, 487, 1, 124, 826, 78, 976, 882, 209, 825, 909, 387, 290, 726, 782, 117, 12, 841, 308, 373, 494, 64, 546, 415, 970, 26, 837, 122, 421, 470, 813, 497, 742, 718, 157, 266, 402, 895, 251, 579, 776, 394, 33, 108, 301, 582, 932, 873, 576, 895, 32, 920, 980, 721, 889, 666, 840, 518, 495, 796, 956, 996, 548, 782, 769, 835, 310, 710, 763, 484, 444, 658, 441, 914, 431, 259, 519, 806, 309, 438, 625, 96, 665, 508, 29, 230, 894, 494, 580, 309, 397, 754, 811, 107, 473, 549, 513, 727, 765, 430, 211, 612, 709, 86, 358, 460, 589, 540, 575, 178, 59, 619, 875, 479, 832, 807, 385, 562, 617, 357, 688, 441, 652, 722, 227, 478, 509, 162, 314, 736, 628, 244, 534, 605, 212, 378, 192, 614, 473, 62, 1, 691, 559, 109, 186, 893, 983, 414, 591, 489, 882, 6, 857, 289, 703, 650, 756, 108, 406, 358, 333, 671, 284, 590, 939, 237, 856, 97, 867, 34, 71, 275, 558, 497, 287, 561, 238, 148, 723, 276, 235, 794, 42, 227, 459, 478, 705, 460, 73, 65, 676, 342, 54, 275, 531, 418, 5, 67, 985, 810, 887, 337, 909, 336, 683, 398, 925, 174, 144, 674, 735, 179, 169, 677, 510, 957, 890, 57, 522, 430, 375, 970, 152, 397, 58, 183, 645, 330, 418, 927, 179, 83, 802, 76, 424, 474, 425, 707, 716, 487, 722, 338, 888, 400, 50, 370, 783, 879, 677, 500, 960, 31, 976, 210, 21, 474, 680, 705, 898, 337, 36, 577, 623, 403, 719, 366, 916, 31, 967, 520, 337, 910, 386, 951, 617, 114, 505, 145, 342, 916, 668, 750, 450, 648, 491, 402, 694, 30, 84, 573, 247, 892, 561, 56, 607, 620, 869, 530, 65, 309, 620, 808, 213, 559, 597, 431, 973, 212, 50, 363, 881, 433, 929, 275, 962, 156, 1000, 804, 191, 478, 100, 390, 263, 544, 117, 392, 423, 204, 499, 784, 327, 540, 866, 570, 788, 429, 778, 13, 5, 489, 112, 893, 36, 553, 531, 69, 581, 448, 165, 198, 916, 57, 81, 79, 603, 338, 875, 772, 562, 823, 361, 284, 290, 897, 888, 633, 611, 117, 689, 825, 358, 88, 383, 842, 2, 439, 401, 144, 798, 302, 573, 43, 557, 104, 318, 713, 950, 56, 734, 628, 695, 712, 158, 207, 398, 307, 49, 337, 651, 492, 189, 286, 164, 925, 813, 184, 471, 199, 195, 259, 336},
		[]int{1, 1, 1, 2, 4, 5, 5, 6, 8, 8, 12, 13, 15, 21, 25, 26, 26, 26, 26, 27, 27, 28, 29, 29, 30, 31, 31, 32, 33, 34, 36, 36, 36, 36, 42, 43, 45, 45, 48, 48, 49, 50, 50, 53, 54, 56, 56, 56, 57, 57, 58, 59, 60, 61, 62, 64, 65, 65, 65, 66, 67, 69, 70, 71, 71, 73, 76, 76, 78, 79, 81, 83, 83, 84, 86, 88, 88, 88, 89, 90, 91, 91, 92, 94, 95, 95, 95, 96, 96, 97, 97, 100, 100, 101, 101, 103, 103, 103, 104, 105, 106, 107, 108, 108, 109, 111, 111, 112, 112, 113, 114, 114, 115, 117, 117, 117, 118, 118, 122, 123, 123, 124, 124, 126, 128, 128, 143, 143, 144, 144, 144, 145, 145, 146, 147, 148, 148, 149, 152, 155, 155, 156, 157, 158, 161, 162, 163, 164, 165, 169, 169, 171, 171, 174, 178, 178, 179, 179, 179, 180, 182, 183, 183, 184, 186, 186, 187, 188, 189, 189, 189, 191, 191, 192, 194, 194, 194, 195, 195, 195, 195, 197, 197, 197, 198, 199, 200, 201, 204, 205, 206, 206, 207, 209, 209, 210, 210, 211, 212, 212, 212, 213, 217, 218, 219, 220, 221, 223, 224, 225, 227, 227, 227, 229, 229, 230, 231, 234, 235, 237, 237, 238, 241, 242, 242, 242, 243, 243, 244, 247, 250, 251, 253, 256, 256, 257, 258, 259, 259, 261, 261, 263, 263, 263, 264, 265, 266, 266, 271, 271, 272, 275, 275, 275, 276, 278, 283, 284, 284, 284, 285, 286, 286, 286, 287, 287, 289, 290, 290, 291, 294, 294, 297, 298, 298, 300, 301, 302, 304, 304, 304, 306, 307, 308, 309, 309, 309, 309, 310, 310, 310, 312, 313, 314, 314, 314, 315, 315, 318, 319, 320, 322, 324, 325, 325, 325, 327, 327, 328, 329, 330, 333, 335, 336, 336, 336, 337, 337, 337, 337, 338, 338, 340, 342, 342, 342, 343, 347, 348, 349, 351, 353, 353, 357, 357, 357, 358, 358, 358, 360, 361, 362, 363, 363, 363, 364, 365, 366, 367, 370, 372, 373, 373, 373, 375, 375, 375, 376, 378, 379, 379, 379, 380, 380, 380, 383, 384, 385, 386, 387, 388, 389, 390, 392, 393, 394, 394, 396, 397, 397, 398, 398, 398, 398, 398, 399, 400, 400, 401, 402, 402, 402, 403, 403, 405, 406, 407, 407, 410, 411, 411, 412, 413, 414, 415, 415, 416, 416, 417, 417, 417, 418, 418, 419, 421, 421, 422, 423, 424, 424, 424, 425, 429, 430, 430, 430, 431, 431, 431, 431, 432, 433, 434, 434, 435, 436, 437, 438, 438, 438, 439, 440, 441, 441, 441, 443, 444, 445, 446, 448, 448, 450, 451, 452, 453, 453, 454, 454, 458, 459, 460, 460, 460, 465, 466, 470, 471, 472, 473, 473, 473, 474, 474, 474, 474, 475, 475, 476, 478, 478, 478, 478, 478, 479, 482, 483, 483, 484, 487, 487, 487, 488, 489, 489, 491, 491, 492, 493, 494, 494, 494, 495, 495, 496, 497, 497, 498, 499, 500, 501, 501, 501, 501, 505, 506, 508, 509, 510, 511, 513, 514, 515, 518, 518, 519, 519, 519, 520, 520, 520, 521, 522, 523, 523, 525, 525, 526, 526, 527, 529, 529, 530, 531, 531, 531, 531, 533, 533, 534, 534, 536, 537, 538, 538, 540, 540, 542, 542, 543, 544, 546, 547, 548, 549, 549, 550, 553, 556, 557, 558, 558, 558, 559, 559, 559, 561, 561, 561, 562, 562, 564, 565, 567, 568, 568, 570, 573, 573, 574, 574, 575, 575, 576, 577, 579, 580, 580, 581, 581, 582, 585, 589, 589, 590, 591, 593, 595, 597, 597, 600, 600, 601, 602, 603, 605, 606, 607, 608, 610, 611, 612, 612, 613, 614, 614, 616, 617, 617, 617, 619, 620, 620, 622, 623, 625, 626, 626, 628, 628, 631, 631, 633, 634, 636, 638, 640, 645, 645, 646, 647, 648, 649, 650, 650, 650, 651, 651, 652, 656, 656, 656, 658, 660, 661, 662, 662, 663, 663, 663, 664, 664, 665, 666, 668, 668, 668, 669, 671, 674, 674, 676, 677, 677, 677, 678, 678, 680, 681, 681, 682, 682, 683, 683, 684, 684, 687, 688, 689, 689, 689, 690, 691, 691, 692, 693, 694, 694, 695, 695, 695, 695, 696, 696, 697, 698, 702, 702, 703, 703, 705, 705, 706, 707, 709, 709, 710, 712, 713, 714, 716, 716, 718, 719, 721, 722, 722, 723, 726, 727, 727, 728, 729, 731, 733, 734, 735, 735, 736, 736, 737, 738, 740, 741, 742, 742, 745, 748, 749, 750, 752, 753, 754, 756, 761, 762, 763, 763, 763, 764, 765, 765, 766, 766, 767, 768, 768, 768, 769, 772, 773, 775, 776, 776, 778, 778, 779, 780, 781, 781, 782, 782, 783, 784, 784, 785, 785, 786, 788, 788, 789, 791, 793, 794, 795, 796, 798, 798, 800, 802, 803, 804, 805, 805, 806, 807, 807, 808, 808, 809, 810, 811, 812, 813, 813, 813, 813, 813, 814, 816, 817, 818, 821, 823, 825, 825, 826, 826, 828, 829, 832, 832, 833, 834, 835, 835, 837, 839, 840, 841, 841, 842, 842, 842, 843, 844, 844, 847, 850, 851, 856, 856, 856, 857, 858, 859, 864, 864, 865, 865, 866, 867, 868, 868, 869, 870, 873, 874, 875, 875, 876, 879, 880, 881, 882, 882, 883, 883, 887, 887, 887, 888, 888, 889, 889, 890, 890, 891, 891, 892, 893, 893, 893, 894, 895, 895, 897, 897, 898, 898, 900, 900, 905, 908, 909, 909, 910, 911, 912, 914, 915, 916, 916, 916, 917, 920, 924, 925, 925, 926, 927, 928, 929, 929, 931, 932, 934, 935, 935, 935, 936, 937, 939, 940, 942, 944, 944, 945, 946, 946, 950, 951, 952, 952, 955, 956, 957, 957, 957, 958, 958, 960, 960, 961, 961, 962, 962, 963, 965, 965, 967, 967, 968, 970, 970, 973, 973, 975, 976, 976, 976, 978, 978, 980, 980, 980, 981, 983, 983, 985, 989, 989, 991, 991, 991, 992, 994, 996, 1000, 1000, 1000},
	},
}

func TestSortArray(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := sortArray(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
