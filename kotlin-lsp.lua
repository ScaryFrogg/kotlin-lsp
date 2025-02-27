local client = vim.lsp.start_client({
	name = "kotlin-lsp",
	cmd = { "/home/vjn/Workspace/kotlin-lsp/kotlin-lsp" },
})

if not client then
	vim.notify("kotlin-lsp: not good")
	return
else
	vim.notify("good")
end
vim.api.nvim_create_autocmd("FileType", {
	pattern = "kotlin",
	callback = function()
		vim.lsp.buf_attach_client(0, client)
	end,
})
