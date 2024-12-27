// Пример данных предметов
const subjects = [
    { id: 1, name: 'Mathematics' },
    { id: 2, name: 'Physics' },
    { id: 3, name: 'Chemistry' },
    { id: 4, name: 'Biology' },
    { id: 5, name: 'History' },
];

// Функция для перехода в аккаунт пользователя
function goToAccount() {
    window.location.href = 'account.html'; // Здесь должен быть путь к странице аккаунта
}

// Функция для перехода на страницу предмета
function goToSubject(subjectId) {
    window.location.href = `resources.html?id=${subjectId}`; // Здесь должен быть путь к странице предмета
}

// Заполнение списка предметов
document.addEventListener('DOMContentLoaded', () => {
    const subjectList = document.getElementById('subjectList');
    subjects.forEach(subject => {
        const li = document.createElement('li');
        li.textContent = subject.name;
        li.onclick = () => goToSubject(subject.id);
        subjectList.appendChild(li);
    });
});