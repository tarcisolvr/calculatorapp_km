document.addEventListener('DOMContentLoaded', function () {
    // Adiciona um evento de clique ao botão de cálculo para garantir que ele execute alguma ação
    const calculateButton = document.querySelector('button.calculate');
    if (calculateButton) {
        calculateButton.addEventListener('click', function (e) {
            e.preventDefault(); // Evita o comportamento padrão do formulário
            document.querySelector('form').submit(); // Submete o formulário
        });
    }

    // Função placeholder para 'Gráfico de Custos'
    window.showGraph = function() {
        alert('Gráfico de Custos será exibido aqui.');
    };
});
