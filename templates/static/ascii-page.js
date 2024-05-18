
const urlParams = new URLSearchParams(window.location.search);
const artText = urlParams.get('artText');

// Находим элемент <pre> по его идентификатору
const preElement = document.getElementById('ascii-page-id');

// Проверяем, есть ли значение artText
if (artText) {
    // Если есть, вставляем его в текст элемента <pre>
    preElement.textContent = artText;
    function downloadAsciiArt() {
        var blob = new Blob([preElement.textContent], { type: 'text/plain' }); // Создаем Blob из текста
        // Создаем временную ссылку для скачивания файла
        var link = document.createElement('a');
        link.href = window.URL.createObjectURL(blob);
        link.download = 'ConvertedText.txt'; // Указываем имя файла для скачивания
        link.click();

        // Очищаем ссылку и освобождаем ресурсы
        window.URL.revokeObjectURL(link.href);
    }

} else {
    // Если нет, выводим сообщение об ошибке
    console.error('Parameter artText is missing in the URL.');
}