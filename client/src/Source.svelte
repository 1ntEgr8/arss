<style>
    .source {
        display: flex;
        justify-content: space-between;
        align-items: center;
        flex-wrap: wrap;
        margin: var(--spacing-small) 0px;
    }

    button {
        min-width: 40px;
        padding: 1px;
    }
</style>

<script>
    import DeleteSource from "./forms/DeleteSource.svelte";
    import EditSource from "./forms/EditSource.svelte";

    export let id;
    export let name = "[no name]";
    export let url = "https://example.com";
    export let noedit = false;
    export let editSource = undefined;
    export let deleteSource = undefined;

    let deleting = false;
    let editing = false;
</script>

<div class="source">
    <a href={url}>
        {name}
    </a>
    {#if !noedit}
    <div>
        <button 
            class={`font-size-xs ${editing ? "highlight" : ""}`}
            on:click={() => editing = true}
            disabled={editing || deleting}
        >edit</button>
        <button 
            class={`font-size-xs ${deleting ? "highlight" : ""}`}
            on:click={() => deleting = true}
            disabled={editing || deleting}
        >del</button>
    </div>
    {/if}
    {#if editing}
        <EditSource
            handleSubmit={(name,url) => {
                editSource(id, {name, url});
                editing = false;
            }}
            handleCancel={() => editing = false}
        />
    {/if}
    {#if deleting}
        <DeleteSource 
            handleYes={() => deleteSource(id)}
            handleNo={() => deleting = false }
        />
    {/if}
</div>
