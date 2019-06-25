$(document).ready(function () {
    new Swiper('.swiper-container', {
        nextButton: '.swiper-button-next',
        prevButton: '.swiper-button-prev',
        loop: true,
        autoplay: true,
        slidesPerView: 3,
        spaceBetween: 10,
        centeredSlides: true,
        disableOnInteraction: false,
        breakpoints: {
            767: {
                nextButton: '.swiper-button-next',
                prevButton: '.swiper-button-prev',
                loop: true,
                autoplay: true,
                slidesPerView: 1,
                spaceBetween: 0,
                centeredSlides: true,
                disableOnInteraction: false,
            }
        }
    });
});