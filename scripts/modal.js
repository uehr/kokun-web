function showResultModal() {
    new Promise(resolve => {
        modalResize()
        resolve($("body").append('<div id="modal-bg"></div>'));
    }).then(res => {
        $("#modal-bg,#modal-main").fadeIn("slow");
    })

    $("#close-modal-button,#modal-bg").click(() => {
        removeResultModal()
    })
}

function modalResize() {
    var w = $(window).width();
    var h = $(window).height();

    var cw = $("#modal-main").outerWidth();
    var ch = $("#modal-main").outerHeight();

    $("#modal-main").css({
        "left": ((w - cw) / 2) + "px",
        "top": ((h - ch) / 2) + "px"
    });
}

function removeResultModal() {
    $("#modal-main,#modal-bg").fadeOut("slow", function () {
        $('#modal-bg').remove();
    });
}