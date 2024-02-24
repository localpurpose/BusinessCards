document.addEventListener("DOMContentLoaded", function() {
    var loader = document.querySelector('.loader');

    if (loader) {

        window.addEventListener('load', function() {
            loader.style.display = 'none';
        });

        window.addEventListener('error', function() {
            loader.style.display = 'none';
        });
    }
});
// swiper
const swiper = new Swiper('.create__slider', {
    effect: 'effect coverflow',
    grabCursor: true,
    loop: false,

    pagination: {
        el: '.swiper-pagination',
        clickable: true,
    },

    // data-color on swiper
    on: {
        slideChange: function () {
            var activeSlide = this.slides[this.activeIndex];
            var colorValue = activeSlide.getAttribute('data-color');

            document.getElementById('selectedCardColorField').value = colorValue;
        },
    },
});

document.addEventListener('DOMContentLoaded', function() {
    const descrInput = document.getElementById('descr');
    const maxChars = 100;

    const container = document.createElement('div');
    container.style.position = 'relative';
    descrInput.parentNode.insertBefore(container, descrInput);
    container.appendChild(descrInput);

    //счетчик символов
    const counter = document.createElement('span');
    counter.textContent = `0/${maxChars}`;
    counter.style.color = '#98bae6';
    counter.style.position = 'absolute';
    counter.style.bottom = '-15px';
    counter.style.right = '50px';


    container.appendChild(counter);

    descrInput.addEventListener('input', function() {
        const inputLength = descrInput.value.length;
        counter.textContent = `${inputLength}/${maxChars}`;
        if (inputLength > maxChars) {
            counter.style.color = 'red';
        } else {
            counter.style.color = '#98bae6';
        }
    });
});
// file input
(function() {
  
    'use strict';
  
    $('.input-file').each(function() {
      var $input = $(this),
          $label = $input.next('.js-labelFile'),
          labelVal = $label.html();
      
     $input.on('change', function(element) {
        var fileName = '';
        if (element.target.value) fileName = element.target.value.split('\\').pop();
        fileName ? $label.addClass('has-file').find('.js-fileName').html(fileName) : $label.removeClass('has-file').html(labelVal);
     });
    });
  
  })();
// inputs container
document.addEventListener('DOMContentLoaded', function () {
const formSteps = document.querySelectorAll('.form-step');
const prevButtons = document.querySelectorAll('.prev');
const nextButtons = document.querySelectorAll('.next');
const form = document.getElementById('form');

$('#tel').inputmask({
    mask: '+9 (999) 999 99 99',
    greedy: false,
    clearIncomplete: true,
    placeholder: '_',
    showMaskOnHover: false,
});

$('#wa').inputmask({
    mask: '9 (999) 999 99 99',
    greedy: false,
    clearIncomplete: true,
    placeholder: '_',
    showMaskOnHover: false,
});

$('#tg').inputmask({
    mask: '@*{*}',
    placeholder: ''
});

$('#url').inputmask({
    mask: 'https://*{*}.{*}/{*}',
    greedy: false,
    skipOptionalPartCharacter: ' ',
    autoUnmask: true,
    definitions: {
      '*': {
        validator: '[0-9A-Za-z-._~:/?#[\\]@!$&\'()*+,;=]',
        cardinality: 1,
        casing: 'lower'
      }
    }
  });


let currentStep = 0;

function showStep(step) {
    formSteps.forEach((formStep, index) => {
        if (index === step) {
            formStep.style.display = 'block';

            if (formStep.classList.contains('templates')) {
                document.querySelector('.dynamic-margin').classList.add('form-step-margin');
                document.querySelector('.form__container').classList.add('container-margin');
            } else {
                document.querySelector('.dynamic-margin').classList.remove('form-step-margin');
                document.querySelector('.form__container').classList.remove('container-margin');
            }
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

function resetErrorMessages() {
    document.querySelectorAll('.error-message').forEach(errorMessage => {
        errorMessage.style.display = 'none';
    });
}

function nextStep(event) {
    event.preventDefault();
    
    const currentFormStep = formSteps[currentStep];
    const requiredInputs = currentFormStep.querySelectorAll('input[required]');
    let allInputsValid = true;

    requiredInputs.forEach(input => {
        if (!input.value.trim()) {
            allInputsValid = false;
            input.style.borderColor = '#5c5cd6';
        } else {
            input.style.borderColor = '';

            if (input.getAttribute('name') === 'name' && !/^[a-zA-Zа-яА-ЯёЁ\s-]{2,}$/u.test(input.value.trim())) {
                allInputsValid = false;
                input.style.borderColor = '#5c5cd6';
            }
            if (input.getAttribute('id') === 'tel' && input.value.trim() && !validator.isMobilePhone(input.value, 'any')) {
                allInputsValid = false;
                input.style.borderColor = '#5c5cd6';
            }

            if (input.getAttribute('type') === 'email' && !validator.isEmail(input.value)) {
                allInputsValid = false;
                input.style.borderColor = '#5c5cd6';
            }
        }
    });

    if (allInputsValid) {
        resetErrorMessages();
        if (currentStep < formSteps.length - 1) {
            currentStep++;
            showStep(currentStep);
        }
    } else {
        document.querySelectorAll('.error-message').forEach(errorMessage => {
            errorMessage.style.display = 'none';
        });
        currentFormStep.querySelector('.error-message').style.display = 'block';
    }
}

function prevStep(event) {
    event.preventDefault();
    if (currentStep > 0) {
        currentStep--;
        showStep(currentStep);
    }
}

nextButtons.forEach(button => button.addEventListener('click', nextStep));
prevButtons.forEach(button => button.addEventListener('click', prevStep));

showStep(currentStep);
});





