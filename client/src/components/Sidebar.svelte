<style>
    .sidebar {
        margin: 0px var(--spacing);
    }

    ul {
        margin: var(--spacing) 0px;
        padding: 0px;
        list-style: none;
    }
    
    li:hover {
        cursor: pointer;
    }

    li:hover:not(.highlight) {
        background: var(--source-hover-color);
    }
</style>

<script>
    import Button from "./Button.svelte";
    import AppBar from "./AppBar.svelte";
    import Source from "./Source.svelte";
    import EditSource from "../forms/EditSource.svelte";

    export let selected = -1;
    export let sources = [];
    export let handleSourceChange;
    export let handleAddSource;
    export let handleDelSource;
    export let handleEditSource;

    let adding = false;
   
</script>

<div class="sidebar">
    <AppBar>
        <span class="font-size-l" slot="main">sources</span>
        <div class="font-size-s" slot="options">
            <button 
                disabled={adding} 
                class={`${adding ? "highlight" : ""}`}
                on:click={() => adding = true}
            >add source</button>
        </div>
    </AppBar>
    
    {#if adding}
        <EditSource
            handleSubmit={async (name, url) => {
                const ok = await handleAddSource(name, url);
                console.log(ok); 
                adding = false;
            }}
            handleCancel={() => adding = false }
        />
    {/if}

    <ul>
        {#each sources as source, i (source.id)}
        <li
            class={i === selected ? "highlight outline" : ""}
        >
            <Source 
                handleClick={() => handleSourceChange(i)}
                selected={i === selected}
                {...source} 
                delSource={handleDelSource} 
                editSource={handleEditSource} 
            />
        </li>
        {/each}
    </ul>
</div>
