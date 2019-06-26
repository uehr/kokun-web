const downloadFileName = "kokun.png"

$(document).ready(function () {
    $("#download-button").click(() => {
        const src = $("#senryu-image").attr("src")
        const blob = dataURIToBlob(src)
        const objUrl = URL.createObjectURL(blob)

        fileDownload(objUrl, downloadFileName)
    })

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

            new Promise(next => {
                next($("#senryu-image-preview").attr("src", "data:image/png;base64," + res.base64Img))
            }).then(next => {
                showResultModal()
            })

        }).catch(err => {
            alert("生成に失敗しました")
        })
    })

});

const dataURIToBlob = (dataURI) => {
    var binStr = atob(dataURI.split(',')[1]),
        len = binStr.length,
        arr = new Uint8Array(len),
        mimeString = dataURI.split(',')[0].split(':')[1].split('')[0]

    for (var i = 0; i < len; i++)
        arr[i] = binStr.charCodeAt(i)

    return new Blob([arr], {
        type: mimeString
    })
}

const fileDownload = (src, name) => {
    const a = document.createElement('a')
    a.href = src
    a.download = name

    a.style.display = 'none'
    document.body.appendChild(a)
    a.click()
    document.body.removeChild(a)
}