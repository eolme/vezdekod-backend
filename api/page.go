package api

import fiber "github.com/gofiber/fiber/v2"

const PAGE_INDEX = `<!DOCTYPE html>
<html>
<body>
  <img id="img" style="height: 400px" />
  <div><pre><code id="info"></code></pre></div>
  <div>
    <button onclick="next()" disabled>skip</button>
    &nbsp;
    &nbsp;
    <button onclick="like()" disabled>like</button>
  </div>
  <script>
    const img = document.getElementById('img');
    const info = document.getElementById('info');
    const buttons = Array.from(document.getElementsByTagName('button'));

    const lock = (state) => buttons.forEach(button => button.disabled = state);

    let meme;

    async function like() {
      lock(true);

      await fetch('/meme/' + meme.id, {
        method: 'POST'
      });

      await next();

      lock(false);
    }

    async function next() {
      lock(true);

      meme = await (await fetch('/meme')).json();
      img.src = meme.photo_url;
      info.textContent = 'Likes: ' + meme.likes + '\nAuthor: ' + meme.author.first_name + ' ' + meme.author.last_name;

      lock(false);
    }

    next()
  </script>
</body>
</html>
`

func SendPageIndex(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "text/html")

	return ctx.Status(200).SendString(PAGE_INDEX)
}

const PAGE_DASHBOARD = `<!DOCTYPE html>
<html>
<head>
  <style>
  table {
    width: 100%;
  }
  table, th, td {
    border: 1px solid #ccc;
    border-collapse: collapse;
  }
  td {
    padding: 4px 8px;
  }
  </style>
</head>
<body>
  <table>
    <thead>
      <tr>
        <th>Place</th>
        <th>Photo</th>
        <th>Likes</th>
        <th>Author</th>
      </tr>
    </thead>
    <tbody></tbody>
  </table>
  <script>
    const tbody = document.getElementsByTagName('tbody')[0];
    const source = new EventSource("/admin/sse");
    source.onmessage = function onmessage(event) {
      let html = '';

      const memes = JSON.parse(event.data);

      memes.forEach((meme, index) => {
        html += '<tr>';

        html += '<td>';
        html += index + 1;
        html += '</td>';

        html += '<td>';
        html += '<a target="_blank" href="' + meme.photo_url + '">' + meme.id + '</a>';
        html += '</td>';

        html += '<td>';
        html += meme.likes;
        html += '</td>';

        html += '<td>';
        html += '<a target="_blank" href="https://vk.com/' + meme.author.screen_name + '">' + meme.author.first_name + ' ' + meme.author.last_name + '</a>';
        html += '</td>';

        html += '</tr>';
      });

      tbody.innerHTML = html;
    }
  </script>
</body>
</html>
`

func SendPageDashboard(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "text/html")

	return ctx.Status(200).SendString(PAGE_DASHBOARD)
}
