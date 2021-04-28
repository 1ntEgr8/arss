<style>
    main {
        display: grid;
        grid-template-columns: 0.2fr 1fr;
        grid-gap: var(--spacing-big);
        margin: 0px var(--spacing-big);
    }
</style>

<script>
    import Header from "./components/Header.svelte";
    import Sidebar from "./components/Sidebar.svelte";
    import Feed from "./components/Feed.svelte";
    
    let sources = [];
    let selected = -1;
    
    async function fetchSources() {
        const res = await fetch(`/sources`) 
        const data = await res.json();
        sources = data;
        if (data?.length > 0) {
            selected = 0;
        }
        return data;
    }

    async function fetchFeed(id) {
        const res = await fetch(`/feed/${id}`);
        const data = await res.json();
        if (res.ok) {
            return data["items"];
        } else {
            throw new Error(data["msg"]);
        }
    }

    function changeSource(id) {
        selected = id;
    }
    
    async function editSource(sourceId, {name, url}) {
        const idx = sources.findIndex(({id}) => sourceId === id)
        sources[idx].name = name;
        sources[idx].url = url;
        
        const res = await fetch(`/sources/edit/${sourceId}`, {
            method: "POST",
            body: JSON.stringify({ name, url })
        });
        if (!res.ok) {
            const data = await res.json();
            throw new Error(data["msg"]);
        }
    }

    async function delSource(sourceId) {
        sources = sources.filter(({id}) => id !== sourceId);
        if (selected >= sources.length) {
            selected = selected - 1;
        }

        const res = await fetch(`/sources/del/${sourceId}`, {
            method: "POST"
        });

        if (!res.ok) {
            const data = await res.json();
            throw new Error(data["msg"]);
        }
    }

    async function addSource(name, url) {
        const res = await fetch(`/sources/add`, {
            method: "POST",
            body: JSON.stringify({name, url})
        });
        const source = await res.json(); 
        
        sources = [...sources, source];
        selected = sources.length - 1;

        if (!res.ok) {
            const data = await res.json();
            throw new Error(data["msg"]);
        }
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
  {:catch}
      Failed to load sources! Check server logs for more info
  {/await}
</main>
