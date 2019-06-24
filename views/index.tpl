<!DOCTYPE html>

<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1">

        <title>kokun</title>

        <link rel="stylesheet" type="text/css" href="../css/index.css">
        <link rel="stylesheet" type="text/css" href="../css/header.css">
        <link rel="stylesheet" href="../css/root.css" />
        <link rel="stylesheet" href="../css/uikit/uikit.min.css" />

        <script src="../scripts/uikit/uikit.min.js"></script>
        <script src="../scripts/uikit/uikit-icons.min.js"></script>
        <script src="../scripts/senryu.js"></script>
        <script src="../scripts/index.js"></script>
    </head>

    <body>
        <div id="site-header">
            <img src="../media/kokun-long.png" id="site-logo">
        </div>

        <div id="contents">
            <form id="test-form">
                <input type="text" id="first-sentence">
                <input type="text" id="second-sentence">
                <input type="text" id="third-sentence">
                <input type="text" id="author-name">
                <button class="uk-button uk-button-default">hogehoge</button>
            </form>
        </div>

        {{.senryuImage}}

</body>
</html>