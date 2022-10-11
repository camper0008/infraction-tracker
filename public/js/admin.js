const remove = async (name) => {
    const apiKey = document.getElementById("api-key");

    const options = { headers: new Headers({ "X-API-KEY": apiKey.value }) };
    const res = await (await fetch(`/api/remove/${name}`, options)).json();
};

const update = async (name) => {
    const apiKey = document.getElementById("api-key");
    const count = document.getElementById("count-" + name);

    const options = { headers: new Headers({ "X-API-KEY": apiKey.value }) };
    const res = await (
        await fetch(`/api/update/${name}/${count.value}`, options)
    ).json();
};

const add = async () => {
    const name = document.getElementById("create-name");
    const apiKey = document.getElementById("api-key");

    const options = { headers: new Headers({ "X-API-KEY": apiKey.value }) };
    const res = await (await fetch(`/api/add/${name.value}`, options)).json();
};
