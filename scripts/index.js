$(document).ready(function () {
    $("#senryu-create-button").click(() => {
        const first = $("#first-sentence").val()
        const second = $("#second-sentence").val()
        const third = $("#third-sentence").val()
        const author = $("#author-name").val()

        getSenryuImage(first, second, third, author).then(res => {
            $("#senryu-image").attr("src", "data:image/png;base64," + res.base64Img)
            $("#first-sentence").val("")
            $("#second-sentence").val("")
            $("#third-sentence").val("")
            $("#author-name").val("")
        })
    })
})