<style>
    .sidebar {
        margin: 0px var(--spacing);
    }

    ul {
        margin: var(--spacing) 0px;
        padding: 0px;
        list-style: none;
    }
</style>

<script>
    import Button from "./Button.svelte";
    import AppBar from "./AppBar.svelte";
    import Source from "./Source.svelte";
    import EditSource from "./forms/EditSource.svelte";

    let sources = [
      { id: 1, name: "lobsters", url: "https://lobste.rs"},
      { id: 2, name: "shtetl-optimized", url: "https://www.scottaaronson.com/blog/"}
    ];
    let adding = false;
    
    function addSource({name, url}) {
        alert("added");
    }

    function editSource(sourceId, {name, url}) {
        // TODO insert api call
        const idx = sources.findIndex(({id}) => sourceId === id)
        sources[idx].name = name;
        sources[idx].url = url;
    }

    function deleteSource(sourceId) {
        // TODO insert api call
        sources = sources.filter(({id}) => id !== sourceId);
    }
</script>

<div class="sidebar">
    <AppBar>
        <span class="font-size-l" slot="main">sources</span>
        <div class="font-size-s" slot="options">
            <button 
                disabled={adding} 
                class={`${adding ? "highlight" : ""}`}
                on:click={() => {console.log("what");adding = true}}
            >add source</button>
        </div>
    </AppBar>
    
    {#if adding}
        <EditSource
            handleSubmit={(name, url) => addSource({name, url})}
            handleCancel={() => adding = false }
        />
    {/if}

    <ul>
        <li>
            <Source id="none" name="all" url="/" noedit/>
        </li>
        {#each sources as source (source.id)}
        <li>
            <Source 
                {...source} 
                {deleteSource} 
                {editSource} 
            />
        </li>
        {/each}
    </ul>
</div>
