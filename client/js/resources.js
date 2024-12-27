// Пример данных разделов
const resources = {
    lectures: ['lecture1.pdf', 'lecture2.pdf'],
    practices: ['practice1.docx'],
    examMaterials: ['exam1.zip'],
    methodical: ['methodical1.pdf', 'methodical2.pdf'],
    practicalWorks: []
};

// Роль пользователя (teacher или student)
const userRole = 'teacher'; // Замените на 'teacher' для тестирования

// Функция для перехода в аккаунт пользователя
function goToAccount() {
    window.location.href = 'account.html'; // Здесь должен быть путь к странице аккаунта
}

function uploadFile(type) {
    const fileInput = document.getElementById(`${type}Upload`);
    if (!fileInput) {
        console.error(`Element with id '${type}Upload' not found`);
        return;
    }
    const file = fileInput.files[0];
    if (file) {
        resources[type].push(file.name);
        appendFileToList(type, file.name);
        fileInput.value = ''; // Очистить input
        alert(`Uploaded ${file.name} to ${type}`);
    } else {
        alert('No file selected');
    }
}

// Функция для скачивания файла
function downloadFile(fileName) {
    // Здесь должен быть код для скачивания файла с сервера
    alert(`Downloading ${fileName}`);
    // Пример запроса:
    // window.location.href = `/download?file=${fileName}`;
}

// Добавление нового файла в список
function appendFileToList(type, fileName) {
    const list = document.getElementById(`${type}List`);
    const div = document.createElement('div');
    div.classList.add('resource-item');
    div.innerHTML = `
        <span>${fileName}</span>
        <button onclick="downloadFile('${fileName}')">⬇</button>
    `;
    list.appendChild(div); // Добавляем новый файл в конец списка
}

function displayResources(type) {
    const list = document.getElementById(`${type}List`);
    if (!list) {
        console.error(`Element with id '${type}List' not found`);
        return;
    }
    list.innerHTML = ''; // Очищаем список перед заполнением
    resources[type].forEach(fileName => {
        appendFileToList(type, fileName);
    });
}

// Инициализация страницы
document.addEventListener('DOMContentLoaded', () => {
    const subjectName = new URLSearchParams(window.location.search).get('subject');
    document.getElementById('subjectName').textContent = subjectName || 'Mathematics';

    // Настройка загрузки файлов для ролей
    if (userRole === 'teacher') {
        document.getElementById('lecturesUpload').style.display = 'block';
        document.getElementById('practicesUpload').style.display = 'block';
        document.getElementById('examMaterialsUpload').style.display = 'block';
        document.getElementById('methodicalsUpload').style.display = 'block';
    } else if (userRole === 'student') {
        document.getElementById('practicalsWorkUpload').style.display = 'block';
    }

    // Отображаем ресурсы для всех типов
    displayResources('lectures');
    displayResources('practices');
    displayResources('examMaterials');
    displayResources('methodical');
    displayResources('practicalWorks');
});