function addToContacts() {
    const nameElement = document.querySelector('.name.contact');
    const organizationElement = document.querySelector('.org.text');
    const phoneElement = document.querySelector('.phone.text');
    const emailElement = document.querySelector('.email.text');
    const websiteElement = document.querySelector('.website.text');

    const name = nameElement ? nameElement.textContent.trim() : '';
    const organization = organizationElement ? organizationElement.textContent.trim() : '';
    const phone = phoneElement ? phoneElement.textContent.trim() : '';
    const email = emailElement ? emailElement.textContent.trim() : '';
    const website = websiteElement ? websiteElement.textContent.trim() : '';

    const vCardText = `
    BEGIN:VCARD
    VERSION:3.0
    N;CHARSET=UTF-8:${name}
    ORG;CHARSET=UTF-8:"${organization}"
    TEL:${phone}
    EMAIL:${email}
    URL:${website}
    END:VCARD
    `;

    const blob = new Blob([vCardText], { type: 'text/vcard;charset=utf-8' });

    const link = document.createElement('a');
    link.href = window.URL.createObjectURL(blob);
    link.setAttribute('download', 'contact.vcf');

    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
}