import axios from "axios";

const title = "Vinland Saga";

const baseUrl = "https://api.mangadex.org";
const resp = await axios({
  method: "GET",
  url: `${baseUrl}/manga`,
  params: {
    title: title,
  },
});

const searchResults = resp.data.data.map((r: any) => ({
  id: r.id,
  title: r.attributes.title.en,
  description: r.attributes.description.en,
  last_chapter: r.attributes.lastChapter,
  last_volume: r.attributes.lastVolume,
  status: r.attributes.status,
}))

console.log(JSON.stringify(searchResults, null, 2));


const mangaId = "5d1fc77e-706a-4fc5-bea8-486c9be0145d";
