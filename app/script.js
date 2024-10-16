// script.js
function searchWeton() {
    // Get the input value
    const inputDate = document.getElementById('inputDate').value;

    // Validate if the input is not empty
    if (!inputDate) {
        document.getElementById('result').innerHTML = "<p class='error'>Pilih tanggal terlebih dahulu.</p>";
        return;
    }

    // Convert from YYYY-MM-DD to DD-MM-YYYY
    const [year, month, day] = inputDate.split("-");
    const formattedDate = `${day}-${month}-${year}`;

    // Clear previous results
    document.getElementById('result').innerHTML = '';

    // Make the AJAX request
    const xhr = new XMLHttpRequest();
    xhr.open("GET", `http://localhost:7723/api/weton/${formattedDate}`, true);
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4 && xhr.status == 200) {
            // Parse the JSON response
            const response = JSON.parse(xhr.responseText);

            // Display the result
            document.getElementById('result').innerHTML = `
        <p>Hari: <strong>${response.data.hari}</strong></p>
        <p>Pasaran: <strong>${response.data.pasaran}</strong></p>
      `;
        } else if (xhr.readyState == 4 && xhr.status != 200) {
            document.getElementById('result').innerHTML = "<p class='error'>Terjadi kesalahan saat mencari weton. Coba lagi nanti.</p>";
        }
    };
    xhr.send();
}

function checkMatchValue() {
    // Get the input value
    var priaDate = document.getElementById('pria-date').value;
    var wanitaDate = document.getElementById('wanita-date').value;

    // Validate if the input is not empty
    if (!priaDate && !wanitaDate) {
        document.getElementById('result').innerHTML = "<p class='error'>Pilih tanggal terlebih dahulu.</p>";
        return;
    }

    // Convert from YYYY-MM-DD to DD-MM-YYYY
    const [year, month, day] = priaDate.split("-");
    const formattedDatePria = `${day}-${month}-${year}`;

    // Convert from YYYY-MM-DD to DD-MM-YYYY
    const [yearWanita, monthWanita, dayWanita] = wanitaDate.split("-");
    const formattedDateWanita = `${dayWanita}-${monthWanita}-${yearWanita}`;

    // Clear previous results
    document.getElementById('result').innerHTML = '';

    // Make the AJAX request
    const xhr = new XMLHttpRequest();
    xhr.open("GET", `http://localhost:7723/api/jodoh?pria=${formattedDatePria}&wanita=${formattedDateWanita}`, true);
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4 && xhr.status == 200) {
            // Parse the JSON response
            const response = JSON.parse(xhr.responseText);

            // Display the result
            document.getElementById('jodoh-result').innerHTML = `
        <p>Tingkat Kecocokan: <strong>${response.data.hasil}</strong></p>
      `;
        } else if (xhr.readyState == 4 && xhr.status != 200) {
            document.getElementById('jodoh-result').innerHTML = "<p class='error'>Terjadi kesalahan saat mencari weton. Coba lagi nanti.</p>";
        }
    };
    xhr.send();
}
