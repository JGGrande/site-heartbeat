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
                Swal.fire({
                    title: "Criado!",
                    text: "Site criado com sucesso!",
                    icon: "success"
                }).then(() => {
                    $modal.classList.remove('active');
                    location.reload();
                });
            } else {
                Swal.fire({
                    title: "Erro!",
                    text: "Erro ao deletar site!",
                    icon: "error"
                });
            }
        } catch (error) {
            console.error('Erro ao enviar requisição:', error);
            Swal.fire({
                title: "Erro!",
                text: "Erro ao deletar site!",
                icon: "error"
            });
        }
    });
})()