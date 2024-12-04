function sweetConfirm(elt, config) {
    Swal.fire(config)
        .then((result) => {
            if (result.isConfirmed) {
                elt.dispatchEvent(new Event('confirmed'));
            }
        });
}

// Alternative, but less clean way, is to not use above JS and do it directly in html:
/*
                    <button type="button" class="bad bg color border"
                        @click="Swal.fire({
                                title: 'Delete these contacts?',
                                showCancelButton: true,
                                confirmButtonText: 'Delete'
                            }).then(result => {
                                if (result.isConfirmed) {
                                    htmx.ajax('DELETE', '/contacts', { source: $root, target: document.body })
                                }
                            });"
                        >Delete
                    </button>
                    */