<style>
    main {
        display: grid;
        grid-template-columns: 0.2fr 1fr;
        grid-gap: var(--spacing-big);
        margin: 0px var(--spacing-big);
    }
</style>

<script>
    import Header from "./Header.svelte";
    import Sidebar from "./Sidebar.svelte";
    import Feed from "./Feed.svelte";

    let sources = [];
    let selected = -1;
    
    async function fetchSources() {
        const res = await fetch("http://localhost:8000/sources") 
        const data = await res.json();
        sources = data;
        if (data?.length > 0) selected = 0;
        return data
    }

    async function fetchFeed(id) {
      const res = await fetch(`http://localhost:8000/feed/${id}`);
      const data = await res.json();
      return data["items"];
    }

    function changeSource(id) {
        selected = id;
    }
    
    async function editSource(sourceId, {name, url}) {
        const idx = sources.findIndex(({id}) => sourceId === id)
        sources[idx].name = name;
        sources[idx].url = url;
        
        const res = await fetch(`http://localhost:8000/sources/edit/${sourceId}`, {
            method: "POST",
            body: JSON.stringify({ name, url })
        });
    }

    async function delSource(sourceId) {
        sources = sources.filter(({id}) => id !== sourceId);
        if (selected >= sources.length) {
            selected = selected - 1;
        }

        const res = await fetch(`http://localhost:8000/sources/del/${sourceId}`, {
            method: "POST"
        });
    }

    async function addSource(name, url) {
        const res = await fetch("http://localhost:8000/sources/add", {
            method: "POST",
            body: JSON.stringify({id: 0, name, url})
        });
        
        sources = [...sources, { id: 7, name, url}];
        selected = sources.length - 1;

        return res.ok;
    }
</script>

<Header />
<main>
  {#await fetchSources()}
        loading sources ...
  {:then}
        <Sidebar 
            {selected} 
            {sources} 
            handleSourceChange={changeSource}
            handleAddSource={addSource}
            handleDelSource={delSource}
            handleEditSource={editSource}
        />
        {#if sources.length > 0}
        <Feed 
            source={sources[selected].name} 
            fetchFeed={fetchFeed(sources[selected].id)}
        />
        {/if}
  {/await}
</main>
