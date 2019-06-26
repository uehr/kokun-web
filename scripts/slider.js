$(document).ready(function () {
    new Swiper('.swiper-container', {
        loop: true,
        autoplay: true,
        slidesPerView: 3,
        spaceBetween: 10,
        centeredSlides: true,
        disableOnInteraction: false,
        navigation: {
            nextEl: '.swiper-button-next',
            prevEl: '.swiper-button-prev',
        },
        breakpoints: {
            767: {
                slidesPerView: 1,
                spaceBetween: 0,
            }
        }
    });
});