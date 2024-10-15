// script.js
function searchWeton() {
    // Get the input value
    const inputDate = document.getElementById('inputDate').value;

    // Validate date input (DD-MM-YYYY format)
    const datePattern = /^\d{2}-\d{2}-\d{4}$/;
    if (!datePattern.test(inputDate)) {
        document.getElementById('result').innerHTML = "<p class='error'>Format tanggal tidak valid. Gunakan format DD-MM-YYYY.</p>";
        return;
    }

    // Clear previous results
    document.getElementById('result').innerHTML = '';

    // Make the AJAX request
    const xhr = new XMLHttpRequest();
    xhr.open("GET", `http://localhost:8080/api/weton/${inputDate}`, true);
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
