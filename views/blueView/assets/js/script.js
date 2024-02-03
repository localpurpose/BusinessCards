function addToContacts() {

    var name = document.querySelector('.name.contact').textContent.trim();
    var organization = document.querySelector('.org.text .org__span').textContent.trim();
    var phone = document.querySelector('.phone.text').textContent.trim();
    var email = document.querySelector('.email.text').textContent.trim();
    var website = document.querySelector('.website.text').textContent.trim();
    var telegram = document.querySelector('.telegram.text').textContent.trim();

    var photoUrl = document.querySelector('.img__container img').getAttribute('src');

    var photoBase64 = '';
    var xhr = new XMLHttpRequest();
    xhr.onload = function () {
        var reader = new FileReader();
        reader.onloadend = function () {
            photoBase64 = reader.result.split(',')[1];
            generateVCard();
        };
        reader.readAsDataURL(xhr.response);
    };
    xhr.open('GET', photoUrl);
    xhr.responseType = 'blob';
    xhr.send();

    function generateVCard() {

        var vCardData = 'BEGIN:VCARD\n' +
            'VERSION:3.0\n' +
            'N:' + name + '\n' +
            'ORG:' + organization + '\n' +
            'TEL:' + phone + '\n' +
            'EMAIL:' + email + '\n' +
            'URL:' + website + '\n' +
            'NOTE:telegram: ' + telegram + '\n' +
            'PHOTO;TYPE=JPEG;ENCODING=BASE64:' + photoBase64 + '\n' +
            'END:VCARD';

        var blob = new Blob([vCardData], { type: 'text/vcard' });
        var url = window.URL.createObjectURL(blob);

        var a = document.createElement('a');
        a.style.display = 'none';
        a.href = url;
        a.download = 'Contact.vcf';
        document.body.appendChild(a);

        a.click();

        document.body.removeChild(a);
        window.URL.revokeObjectURL(url);
    }
}
