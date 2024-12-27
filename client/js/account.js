// Пример данных пользователя и предметов
const user = {
    name: 'Sabina',
    studyGroup: 'IKBO-01-22',
    subjects: ['Mathematics', 'Physics', 'Chemistry', 'Biology', 'History']
};

// Заполнение информации об аккаунте
document.addEventListener('DOMContentLoaded', () => {
    document.getElementById('userName').textContent = user.name;
    document.getElementById('studyGroup').textContent = user.studyGroup;

    const subjectList = document.getElementById('subjectList');
    user.subjects.forEach(subject => {
        const li = document.createElement('li');
        li.textContent = subject;
        subjectList.appendChild(li);
    });
});

// Функция для перехода на предыдущую страницу
function goBack() {
    window.history.back();
}

// Функция для выхода из аккаунта
function logout() {
    // Здесь должен быть код для выхода из аккаунта, например, очистка токенов
    alert('Logged out successfully!');
    window.location.href = 'login.html'; // Переход на страницу входа
}
