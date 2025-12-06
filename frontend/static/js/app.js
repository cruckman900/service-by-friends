async function loadPartial(id, file) {
  const response = await fetch(`/templates/partials/${file}`);
  document.getElementById(id).innerHTML = await response.text();
}

async function loadContent(page) {
  const response = await fetch(`/templates/content/${page}`);
  document.getElementById("content").innerHTML = await response.text();
}

// Sidebar toggle
function toggleSidebar() {
  document.getElementById("sidebar").classList.toggle("collapsed");
}

// Accordian toggle
function toggleAccordion(element) {
  const content = element.nextElementSibling;
  content.style.display = content.style.display === "block" ? "none" : "block";
}

async function buildSidebar() {
  const response = await fetch("/data/services.json");
  const services = await response.json();
  const accordion = document.getElementById("accordion");

  accordion.innerHTML = "";

  Object.keys(services).forEach((category) => {
    const section = document.createElement("div");

    // Header with icon + label
    const header = document.createElement("h3");
    header.innerHTML = `<span class="icon">${services[category][0].icon}</span>
                        <span class="label">${category}</span>`;
    header.onclick = () => toggleAccordion(header);

    const content = document.createElement("div");
    content.classList.add("accordion-content");

    services[category].forEach((item) => {
      const link = document.createElement("a");
      link.href = `/templates/content/${item.page}`;
      link.textContent = item.name;
      link.onclick = (e) => {
        e.preventDefault();
        loadContent(item.page);
      };
      content.appendChild(link);
    });

    section.appendChild(header);
    section.appendChild(content);
    accordion.appendChild(section);
  });
}

function wireHeaderLinks() {
  document.querySelectorAll("#header nav a").forEach((link) => {
    link.addEventListener("click", function (e) {
      e.preventDefault();
      loadContent(this.dataset.page);
    });
  });
}

// Inititalize app
document.addEventListener("DOMContentLoaded", async () => {
  // Load partials
  await loadPartial("header", "header.html");
  wireHeaderLinks(); // attach handlers here
  await loadPartial("sidebar-container", "sidebar.html");
  await loadPartial("footer", "footer.html");
  await buildSidebar();

  // Load default content (home page)
  await loadContent("home.html");
});
