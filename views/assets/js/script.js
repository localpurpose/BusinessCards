window.addEventListener('resize', function() {
    var imgContainer = document.querySelector('.img__container');
    var windowHeight = window.innerHeight;
    var targetHeight = 5 * windowHeight; // 40% высоты экрана
    imgContainer.style.height = targetHeight + 'px';
});

// Вызовем событие resize после загрузки страницы
window.dispatchEvent(new Event('resize'));


// Открытие модального окна
document.querySelector('.img').addEventListener('click', function() {
    document.getElementById('myModal').style.display = 'block';
    document.getElementById('modalImg').src = 'assets/img/frame.png';
  });
  
  // Закрытие модального окна
  document.querySelector('.close').addEventListener('click', function() {
    document.getElementById('myModal').style.display = 'none';
  });


// анимация печати  
document.addEventListener('DOMContentLoaded', function () {
    const spans = document.querySelectorAll('.typing-animation');
    spans.forEach((span, index) => {
        const text = span.textContent.trim();
        span.textContent = '';
        setTimeout(() => typeText(span, text), index * 1000); // Задержка в секундах
    });
});

function typeText(element, text) {
    let index = 0;

    function type() {
        element.textContent += text[index];
        index++;

        if (index < text.length) {
            setTimeout(type, 40); // Интервал между символами
        }
    }

    type();
}
