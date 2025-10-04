function daysLeft(entryDate, visaExpiry, allowedDays, stayType) {
  const now = new Date();
  if (stayType === 'visa' && visaExpiry) {
    const expiry = new Date(visaExpiry);
    return Math.floor((expiry - now) / (1000 * 60 * 60 * 24));
  } else if (stayType === 'visa-free' && allowedDays) {
    const deadline = new Date(entryDate);
    deadline.setDate(deadline.getDate() + parseInt(allowedDays));
    return Math.floor((deadline - now) / (1000 * 60 * 60 * 24));
  }
  return null;
}

function fetchAndRender(url) {
  fetch(url)
    .then(res => res.json())
    .then(data => {
      const tbody = document.querySelector("#migrant-table tbody");
      tbody.innerHTML = "";
      data.forEach(m => {
        const row = document.createElement("tr");
        const left = daysLeft(m.entry_date, m.visa_expiry, m.allowed_days, m.stay_type);
        row.className = (left !== null && left < 0) ? "expired" : "valid";
        row.innerHTML = `
          <td>${m.full_name}</td>
          <td>${m.passport}</td>
          <td>${m.nationality}</td>
          <td>${m.stay_type}</td>
          <td>${m.entry_date?.slice(0, 10)}</td>
          <td>${m.visa_expiry?.slice(0, 10) || "-"}</td>
          <td>${left !== null ? left + " дн." : "-"}</td>
          ${url.includes('expired') ? '' : `<td><button onclick="deleteMigrant(${m.id})">Удалить</button></td>`}
        `;
        tbody.appendChild(row);
      });
    })
    .catch(error => {
      console.error("Ошибка при получении данных:", error);
    });
}

function deleteMigrant(id) {
  if (!confirm('Удалить этого мигранта?')) return;
  fetch(`/api/migrants/${id}`, { method: 'DELETE' })
    .then(res => {
      if (res.ok) location.reload();
      else alert('Ошибка удаления');
    });
}