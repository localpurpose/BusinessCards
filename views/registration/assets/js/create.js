window.addEventListener('load', function () {

    var loader = document.getElementById('loader');
    loader.style.display = 'none';


});

document.addEventListener('DOMContentLoaded', function () {
const formSteps = document.querySelectorAll('.form-step');
const prevButtons = document.querySelectorAll('.prev');
const nextButtons = document.querySelectorAll('.next');
const form = document.getElementById('form');

let currentStep = 0;

function showStep(step) {
    formSteps.forEach((formStep, index) => {
        if (index === step) {
            formStep.style.display = 'block';
        } else {
            formStep.style.display = 'none';
        }
    });

    if (step === 0) {
        prevButtons.forEach(prevButton => prevButton.style.display = 'none');
    } else {
        prevButtons.forEach(prevButton => prevButton.style.display = 'inline-block');
    }

    if (step === formSteps.length - 1) {
        nextButtons.forEach(nextButton => nextButton.style.display = 'none');
    } else {
        nextButtons.forEach(nextButton => nextButton.style.display = 'inline-block');
    }

    const progress = ((step + 1) / formSteps.length) * 100;
    document.querySelector('.progress-bar').style.width = progress + '%';
    document.querySelector('.progress-bar').setAttribute('aria-valuenow', progress);
}

function nextStep(event) {
    event.preventDefault(); // Prevent default form submission
    if (currentStep < formSteps.length - 1) {
        currentStep++;
        showStep(currentStep);
    }
}

function prevStep(event) {
    event.preventDefault(); // Prevent default form submission
    if (currentStep > 0) {
        currentStep--;
        showStep(currentStep);
    }
}

nextButtons.forEach(button => button.addEventListener('click', nextStep));
prevButtons.forEach(button => button.addEventListener('click', prevStep));

showStep(currentStep);

const swiper = new Swiper('.create__slider', {
    effect: 'effect coverflow',
    grabCursor: true,
    loop: false,

    pagination: {
        el: '.swiper-pagination',
        clickable: true,
    },

    on: {
        slideChange: function () {
            var activeSlide = this.slides[this.activeIndex];
            var colorValue = activeSlide.getAttribute('data-color');

            document.getElementById('selectedCardColorField').value = colorValue;
        },
    },
});

});

