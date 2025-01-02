"use strict";

(function(){
    const $modal = document.querySelector('#site-modal');
    const $openModalButton = document.querySelector('#open-modal-button');
    const $closeModalButton = document.querySelector('#close-modal-button');
    const $submitButton = document.querySelector('#submit-button');

    $openModalButton.addEventListener('click', () => {
        $modal.classList.add('active');
    });

    $closeModalButton.addEventListener('click', () => {
        $modal.classList.remove('active');
    });

    $submitButton.addEventListener('click', async () => {
        const name = document.querySelector('#site-name').value;
        const url = document.querySelector('#site-url').value;

        if (!name || !url) {
            alert('Preencha todos os campos!');
            return;
        }

        try {
            const protocol = window.location.protocol;
            const port = window.location.port || "80";
            const apiUrl = `${protocol}//${window.location.hostname}:${port}/criar-monitoramento`;

            const response = await fetch(apiUrl, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ nome: name, url: url }),
            });

            if (response.status === 201) {
                alert('Site adicionado com sucesso!');
                $modal.classList.remove('active');
                location.reload();
            } else {
                alert('Ocorreu um erro ao adicionar o site.');
            }
        } catch (error) {
            console.error('Erro ao enviar requisição:', error);
            alert('Erro ao conectar-se ao servidor.');
        }
    });
})()