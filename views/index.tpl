<!DOCTYPE html>

<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link rel="icon" type="image/x-icon" href="../media/favicon.ico">

        <title>kokun</title>

        <link rel="stylesheet" type="text/css" href="../css/index.css">
        <link rel="stylesheet" type="text/css" href="../css/header.css">
        <link rel="stylesheet" type="text/css" href="../css/root.css" />
        <link rel="stylesheet" type="text/css" href="../css/swiper/swiper.css" />

        <script src="https://code.jquery.com/jquery-3.3.1.js"></script>
        <script src="../scripts/senryu.js"></script>
        <script src="../scripts/index.js"></script>
        <script src="../scripts/slider.js"></script>
        <script src="../scripts/swiper/swiper.min.js"></script>
    </head>

    <body>
        <div id="site-header">
            <img src="../media/kokun-long.png" id="site-logo">
        </div>

        <div id="contents">

            <!-- Swiper START -->
            <div class="swiper-container">
                <h1 id="site-description" class="uk-heading-line uk-text-center"><span>あなたの<span class="theme-color">俳句</span>/<span class="theme-color">川柳</span>を本格的に</span></h1>

	            <!-- メイン表示部分 -->
	            <div class="swiper-wrapper">
		            <div class="swiper-slide">
                        <img src="../media/examples/1.png" class="example-image">
                    </div>

		            <div class="swiper-slide">
                        <img src="../media/examples/2.png" class="example-image">
                    </div>

		            <div class="swiper-slide">
                        <img src="../media/examples/3.png" class="example-image">
                    </div>

		            <div class="swiper-slide">
                        <img src="../media/examples/4.png" class="example-image">
                    </div>

		            <div class="swiper-slide">
                        <img src="../media/examples/5.png" class="example-image">
                    </div>

	            </div>

                <div class="swiper-button-prev swiper-button-white"></div>
                <div class="swiper-button-next swiper-button-white"></div>

            </div>

            <!-- Swiper END -->
            <div id="senryu-form">
                <div >
                    <input id="first-sentence" type="text" placeholder="五">
                    <input id="second-sentence" type="text" placeholder="七">
                    <input id="third-sentence" type="text" placeholder="五">
                    <input id="author-name" type="text" placeholder="名前">
                </div>

                <button id="senryu-create-button">作成</button>
            </form>

            <div id="result-popup">
                <img id="senryu-image" src="">
            </div>

        </div>

</body>
</html>