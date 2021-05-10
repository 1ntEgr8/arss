<style>
    .source {
        display: flex;
        justify-content: space-between;
        align-items: center;
        flex-wrap: wrap;
        padding: var(--spacing-small) var(--spacing-small);
    }

    button {
        min-width: 40px;
        padding: 1px;
        cursor: pointer;
    }

    button:hover:not(:disabled) {
        background: #d5d458;
    }
</style>

<script>
    import DeleteSource from "../forms/DeleteSource.svelte";
    import EditSource from "../forms/EditSource.svelte";

    export let id;
    export let name = "[no name]";
    export let url = "[no url]";
    export let noedit = false;
    export let editSource;
    export let delSource;
    export let handleClick;
    export let selected = false;

    let deleting = false;
    let editing = false;
</script>

<div 
    class="source"
    on:click={handleClick}
    >
    {name}
    {#if !noedit && selected}
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

        {#if editing}
            <EditSource
                handleSubmit={(name,url) => {
                    editSource(id, {name, url});
                    editing = false;
                }}
                handleCancel={() => editing = false}
                name={name}
                url={url}
            />
        {/if}
        {#if deleting}
            <DeleteSource 
                handleYes={() => delSource(id)}
                handleNo={() => deleting = false }
            />
        {/if}
    {/if}
</div>
