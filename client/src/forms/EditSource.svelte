<style>
    .form {
        width: 100%;
        margin: var(--spacing-small) 0px;
        display: grid;
        grid-template-columns: repeat(4,1fr);
        grid-gap: var(--spacing-small);
    }
    
    .btn-container {
        display: grid;
        grid-template-columns: repeat(2,1fr);
        grid-column: 1/5;
        grid-gap: var(--spacing-small);
    }

    label {
        align-self: center;
    }

    button {
        font-size: var(--font-size-s);
    }

    #cancel {
        background: var(--color-danger);
    }

    label {
        grid-column: 1/2; 
    }

    input {
        grid-column: 2/5;
    }
</style>

<script>
    export let handleSubmit;
    export let handleCancel;
    let name;
    let url;
</script>

<form 
    class="form font-size-m" 
    on:submit|preventDefault={() => handleSubmit(name, url)}
>
    <label for="name">
        name
    </label>
    <input required id="name" bind:value={name}>

    <label for="url">
        url
    </label>
    <input required id="url" bind:value={url} type="url">
    
    <div class="btn-container">
        <button type="submit">done</button>
        <button 
            id="cancel" 
            type="button" 
            on:click={handleCancel}
        >cancel</button>
    </div>
</form>
