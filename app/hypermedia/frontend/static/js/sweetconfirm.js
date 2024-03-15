// couldn't get this to work
/*
function sweetConfirm(elt, config) {
    Swal.fire(config)
        .then((result) => {
            if (result.isConfirmed) {
                elt.dispatchEvent(new Event('confirmed'));
            }
        });
}

*/
/*                    <button type="button" class="bad bg color border"
                            hx-delete="/contacts" hx-target="body" hx-trigger="confirmed"
                            @click="sweetConfirm($el,
                                    { title: 'Delete these contacts?',
                                      showCancelButton: true,
                                      confirmButtonText: 'Delete'})">
                                      Delete
                    </button>*/