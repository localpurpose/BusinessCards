function addToContacts() {
    // Получаем информацию из полей с классом .contact
    const name = document.querySelector('.name.contact').textContent.trim();
    const organization = document.querySelector('.org.text').textContent.trim();
    const phone = document.querySelector('.phone.text').textContent.trim();
    const email = document.querySelector('.email.text').textContent.trim();
    const website = document.querySelector('.website.text').textContent.trim();
    const telegram = document.querySelector('.telegram.text').textContent.trim();
    const whatsapp = document.querySelector('.whatsapp.text').textContent.trim();

    // Формируем текст для контакта
    const contactText = `Имя: ${name}\nОрганизация: ${organization}\nТелефон: ${phone}\nПочта: ${email}\nВебсайт: ${website}\nTelegram: ${telegram}\nWhatsApp: ${whatsapp}`;

    // Проверяем поддержку Web Share API в браузере
    if (navigator.share) {
        navigator.share({
            title: 'Контактная информация',
            text: contactText,
        })
        .then(() => console.log('Контакт добавлен успешно'))
        .catch((error) => console.error('Ошибка при добавлении контакта', error));
    } else {
        // В случае, если Web Share API не поддерживается, можно предоставить альтернативный опыт для пользователя
        console.log('Web Share API не поддерживается в этом браузере');
        // Тут можно предложить пользователю скопировать текст вручную или использовать другие методы добавления контакта.
    }
}