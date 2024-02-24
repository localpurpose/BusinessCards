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
