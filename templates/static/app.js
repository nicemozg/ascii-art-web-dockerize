
var dataText = ""
    // Функция для отправки формы
    function submitForm() {
      var text = document.getElementById("textArea").value;
        var selectElement = document.getElementById("banner");
        var selectedOption = selectElement.value;

        var requestData = {
        text: text,
        banner: selectedOption
      };
      fetch('/ascii-art', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(requestData)
      })
        .then(response => {
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          return response.text();
        })
        .then(data => {
          // Обработка успешного ответа от сервера
          var preElement = document.getElementById('ascii-art');
          preElement.textContent = data;
            dataText = data
          console.log(data);
        })
        .catch(error => {
            console.error('There has been a problem with your fetch operation:', error);
            alert('Server does not response!!!');
        });
    }

function goToAsciiPage() {
    // Сохраняем значение data в переменной artText
    // Формируем URL для перехода на страницу ascii-art-page.html с параметром artText
    var url = '/ascii-art-page.html?artText=' + encodeURIComponent(dataText);
    // Выполняем переход на новую страницу
    window.location.href = url;
}

