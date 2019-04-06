/**
 * jQuery pagination plugin v1.0
 * http://esimakin.github.io/twbs-pagination/
 *
 * Copyright 2014, Eugene Simakin
 * Released under Apache 2.0 license
 * http://apache.org/licenses/LICENSE-2.0.html
 */
;
(function ($, window, document, undefined) {

    'use strict';

    var old = $.fn.twbsPagination;

    // PROTOTYPE AND CONSTRUCTOR

    var TwbsPagination = function (element, options) {
        this.$element = $(element);
        this.options = $.extend({}, $.fn.twbsPagination.defaults, options);
        this.init(this.options);
    };

    TwbsPagination.prototype = {
        constructor: TwbsPagination,
        init: function (options) {
            this.options = options;
            this.currentPages = this.getPages(this.options.startPage);
            if (this.options.startPage < 1) {
                this.options.startPage = 1;
            }
            if (this.options.totalPage <= 0) {
                this.options.totalPage = 1;
            }
            if (!$.isNumeric(this.options.visiblePage) && !this.options.visiblePage) {
                this.options.visiblePage = this.options.totalPage;
            }
            if (this.options.onPageClick instanceof Function) {
                this.$element.bind('page', this.options.onPageClick);
            }
            var tagName = (typeof this.$element.prop === 'function') ?
                    this.$element.prop('tagName') : this.$element.attr('tagName');
            if (tagName === 'UL') {
                this.$listContainer = this.$element;
                this.$listContainer.html("");
            } else {
                this.$listContainer = $('<ul></ul>');
            }
            this.$listContainer.addClass(this.options.paginationClass);
            this.$listContainer.append(this.buildListItems(this.currentPages.numeric));
            if (tagName !== 'UL') {
                this.$element.append(this.$listContainer);
            }
            this.show(this.options.startPage);
        },
        show: function (page) {
            if (page < 1 || page > this.options.totalPage) {
                throw new Error('Page is incorrect.');
            }
            this.render(this.getPages(page));
            this.setupEvents();
            this.$element.trigger('page', page);
        },
        buildListItems: function (pages) {
            var $listItems = $();

            if (this.options.first) {
                $listItems = $listItems.add(this.buildItem('first', 1));
            }

            if (this.options.prev) {
                $listItems = $listItems.add(this.buildItem('prev', 1));
            }

            for (var i = 0; i < pages.length; i++) {
                $listItems = $listItems.add(this.buildItem('page', pages[i]));
            }

            if (this.options.next) {
                $listItems = $listItems.add(this.buildItem('next', 2));
            }

            if (this.options.last) {
                $listItems = $listItems.add(this.buildItem('last', this.options.totalPage));
            }

            return $listItems;
        },
        buildItem: function (type, page) {
            var itemContainer = $('<li></li>'),
                    itemContent = $('<a></a>'),
                    itemText = null;
            itemContainer.addClass(type);
            itemContainer.attr('data-page', page);
            switch (type) {
                case 'page':
                    itemText = page;
                    break;
                case 'first':
                    itemText = this.options.first;
                    break;
                case 'prev':
                    itemText = this.options.prev;
                    break;
                case 'next':
                    itemText = this.options.next;
                    break;
                case 'last':
                    itemText = this.options.last;
                    break;
                default:
                    break;
            }
            itemContainer.append(itemContent.attr('href', this.href(page)).text(itemText));
            return itemContainer;
        },
        getPages: function (currentPage) {
            var pages = [];
            var half = Math.floor(this.options.visiblePage / 2);
            var start = currentPage - half + 1 - this.options.visiblePage % 2;
            var end = currentPage + half;
            if (start <= 0) {
                start = 1;
                end = this.options.visiblePage;
            }
            if (end > this.options.totalPage) {
                start = this.options.totalPage - this.options.visiblePage + 1;
                end = this.options.totalPage;
            }
            if (start <= 0) {
                start = 1;
            }
            var itPage = start;
            while (itPage <= end) {
                pages.push(itPage);
                itPage++;
            }
            return {"currentPage": currentPage, "numeric": pages};
        },
        render: function (pages) {
            if (!this.equals(this.currentPages.numeric, pages.numeric)) {
                this.$listContainer.children().remove();
                this.$listContainer.append(this.buildListItems(pages.numeric));
                this.currentPages = pages;
            }

            this.$listContainer.find('.page').removeClass('active');
            this.$listContainer.find('.page').filter('[data-page="' + pages.currentPage + '"]').addClass('active');

            this.$listContainer.find('.first')
                    .toggleClass('disabled', pages.currentPage === 1);

            this.$listContainer.find('.last')
                    .toggleClass('disabled', pages.currentPage === this.options.totalPage);

            var prev = pages.currentPage - 1;
            this.$listContainer.find('.prev')
                    .toggleClass('disabled', pages.currentPage === 1)
                    .data('page', prev > 1 ? prev : 1);

            var next = pages.currentPage + 1;
            this.$listContainer.find('.next')
                    .toggleClass('disabled', pages.currentPage === this.options.totalPage)
                    .data('page', next < this.options.totalPage ? next : this.options.totalPage);
        },
        setupEvents: function () {
            var base = this;
            this.$listContainer.find('li').each(function () {
                var $this = $(this);
                $this.off();
                if ($this.hasClass('disabled') || $this.hasClass('active'))
                    return;
                $this.click(function () {
                    base.show(parseInt($this.data('page'), 10));
                });
            });
        },
        equals: function (oldArray, newArray) {
            var i = 0;
            while ((i < oldArray.length) || (i < newArray.length)) {
                if (oldArray[i] !== newArray[i]) {
                    return false;
                }
                i++;
            }
            return true;
        },
        href: function (c) {
            return this.options.href.replace(this.options.hrefVariable, c);
        }
    };

    // PLUGIN DEFINITION

    $.fn.twbsPagination = function (option) {
        var args = Array.prototype.slice.call(arguments, 1);
        var methodReturn;

        var $set = this.each(function () {
            var $this = $(this);
            var data = $this.data('twbs-pagination');
            var options = typeof option === 'object' && option;

            if (!data)
                $this.data('twbs-pagination', (data = new TwbsPagination(this, options)));
            if (typeof option === 'string')
                methodReturn = data[ option ].apply(data, args);
        });

        return (methodReturn === undefined) ? $set : methodReturn;
    };

    $.fn.twbsPagination.defaults = {
        totalPage: 0,
        startPage: 1,
        visiblePage: 5,
        href: 'javascript:void(0);',
        hrefVariable: '{{number}}',
        first: '首页',
        prev: '上页',
        next: '下页',
        last: '尾页',
        paginationClass: 'pagination-sm pagination',
        onPageClick: null,
        version: '1.1'
    };

    $.fn.twbsPagination.Constructor = TwbsPagination;

    $.fn.twbsPagination.noConflict = function () {
        $.fn.twbsPagination = old;
        return this;
    };

})(jQuery, window, document);