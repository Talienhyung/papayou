document.addEventListener('DOMContentLoaded', function () {
    const form = document.getElementById('players-form');
    const inputsContainer = document.getElementById('inputs-container');
    const addButton = document.getElementById('add-player');
    const submitButton = document.getElementById('submit-players');

    let playerCount = 2; // Initial player count

    addButton.addEventListener('click', function () {
        if (playerCount < 8) {
            playerCount++;
            const newPlayerInput = document.createElement('div');
            newPlayerInput.classList.add('player-input');
            newPlayerInput.innerHTML = `<label for="player${playerCount}">Joueur ${playerCount}:</label><br>
                                        <input type="text" id="player${playerCount}" name="player${playerCount}">`;
            inputsContainer.appendChild(newPlayerInput);
        }
    });

    submitButton.addEventListener('click', function () {
        const playerNames = [];
        for (let i = 1; i <= playerCount; i++) {
            const playerName = document.getElementById(`player${i}`).value;
            playerNames.push(playerName);
        }
    });
});
